import { authActions, type User } from '$lib/stores/authStore';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';
console.log('AuthService API_BASE_URL:', API_BASE_URL);

export interface LoginCredentials {
    username: string;
    password: string;
}

export interface RegisterCredentials {
    username: string;
    password: string;
}

export interface AuthResponse {
    token?: string;
    data?: {
        user: User;
    };
    error?: string;
}

class AuthService {
    private baseURL: string;

    constructor(baseURL: string) {
        this.baseURL = baseURL;
    }

    // Generic request method
    private async request<T>(
        endpoint: string,
        options: RequestInit = {}
    ): Promise<T> {
        const url = `${this.baseURL}${endpoint}`;

        try {
            const response = await fetch(url, {
                ...options,
                headers: {
                    'Content-Type': 'application/json',
                    ...options.headers,
                },
            });

            const data = await response.json();

            if (!response.ok) {
                throw new Error(data.error || `HTTP ${response.status}: ${response.statusText}`);
            }

            return data;
        } catch (error) {
            const errorMessage = error instanceof Error ? error.message : 'Unknown error occurred';
            throw new Error(errorMessage);
        }
    }

    // Login user
    async login(credentials: LoginCredentials): Promise<{ token: string; user: User }> {
        authActions.setLoading(true);
        authActions.clearError();

        try {
            // First, login to get token
            const loginResponse = await this.request<{ token: string }>('/auth/login', {
                method: 'POST',
                body: JSON.stringify(credentials),
            });

            if (!loginResponse.token) {
                throw new Error('No token received from server');
            }

            // Then get user profile using the token
            const profileResponse = await this.request<{ user: User }>('/user/profile', {
                method: 'GET',
                headers: {
                    'Authorization': `Bearer ${loginResponse.token}`,
                },
            });

            const authData = {
                token: loginResponse.token,
                user: profileResponse.user
            };

            // Store auth data
            authActions.setAuth(authData.token, authData.user);

            return authData;
        } catch (error) {
            const errorMessage = error instanceof Error ? error.message : 'Login failed';
            authActions.setError(errorMessage);
            throw error;
        } finally {
            authActions.setLoading(false);
        }
    }

    // Register user
    async register(credentials: RegisterCredentials): Promise<{ user: User }> {
        authActions.setLoading(true);
        authActions.clearError();

        try {
            const response = await this.request<{ data: User }>('/auth/signup', {
                method: 'POST',
                body: JSON.stringify(credentials),
            });

            return { user: response.data };
        } catch (error) {
            const errorMessage = error instanceof Error ? error.message : 'Registration failed';
            authActions.setError(errorMessage);
            throw error;
        } finally {
            authActions.setLoading(false);
        }
    }

    // Logout user
    logout() {
        authActions.logout();
    }

    // Get user profile (to verify token is still valid)
    async getProfile(token: string): Promise<User> {
        const response = await this.request<{ user: User }>('/user/profile', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`,
            },
        });

        return response.user;
    }
}

// Export singleton instance
export const authService = new AuthService(API_BASE_URL);
