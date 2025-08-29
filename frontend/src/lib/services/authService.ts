import { authActions, type User } from '$lib/stores/authStore';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';
console.log('AuthService API_BASE_URL:', API_BASE_URL);

export interface LoginCredentials {
    username_or_email: string;
    password: string;
}

export interface RegisterCredentials {
    username: string;
    email: string;
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

    // Custom HTTP error to propagate status and payload
    private buildError(response: Response, payload: any): Error {
        const err: any = new Error(payload?.error || `HTTP ${response.status}: ${response.statusText}`);
        err.status = response.status;
        err.payload = payload;
        return err as Error & { status?: number; payload?: any };
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

            const text = await response.text();
            const data = text ? JSON.parse(text) : {};

            if (!response.ok) {
                throw this.buildError(response, data);
            }

            return data;
        } catch (error) {
            if (error instanceof Error) throw error;
            throw new Error('Unknown error occurred');
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
        } catch (error: unknown) {
            const err = error as Error & { status?: number; payload?: any };
            // Friendly message if backend says needs verification
            if ((err as any)?.status === 401 && (err as any)?.payload?.needs_verification) {
                authActions.setError('Please verify your email address before logging in.');
                throw err;
            }
            const errorMessage = err?.message || 'Login failed';
            authActions.setError(errorMessage);
            throw err;
        } finally {
            authActions.setLoading(false);
        }
    }

    // Register user
    async register(credentials: RegisterCredentials): Promise<{ message: string; user_id?: number }> {
        authActions.setLoading(true);
        authActions.clearError();

        try {
            const response = await this.request<{ message: string; user_id?: number }>('/auth/signup', {
                method: 'POST',
                body: JSON.stringify(credentials),
            });
            return { message: response.message, user_id: response.user_id };
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

    // Verify email using token
    async verifyEmail(token: string): Promise<{ message: string }> {
        return this.request<{ message: string }>(`/auth/verify-email?token=${encodeURIComponent(token)}`, {
            method: 'GET',
        });
    }

    // Resend verification email
    async resendVerification(email: string): Promise<{ message: string }> {
        return this.request<{ message: string }>(`/auth/resend-verification`, {
            method: 'POST',
            body: JSON.stringify({ email })
        });
    }
}

// Export singleton instance
export const authService = new AuthService(API_BASE_URL);
