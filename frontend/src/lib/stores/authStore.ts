import { writable } from 'svelte/store';
import { browser } from '$app/environment';

export interface User {
    id: number;
    username: string;
}

export interface AuthState {
    user: User | null;
    token: string | null;
    isAuthenticated: boolean;
    isLoading: boolean;
    error: string | null;
}

const initialState: AuthState = {
    user: null,
    token: null,
    isAuthenticated: false,
    isLoading: false,
    error: null
};

// Create the auth store
export const authStore = writable<AuthState>(initialState);

// Auth actions
export const authActions = {
    // Initialize auth state from localStorage on app start
    init() {
        if (!browser) return;
        
        const token = localStorage.getItem('auth_token');
        const userStr = localStorage.getItem('auth_user');
        
        if (token && userStr) {
            try {
                const user = JSON.parse(userStr);
                authStore.update(state => ({
                    ...state,
                    token,
                    user,
                    isAuthenticated: true
                }));
            } catch (error) {
                // Clear invalid data
                localStorage.removeItem('auth_token');
                localStorage.removeItem('auth_user');
            }
        }
    },

    // Set authentication data
    setAuth(token: string, user: User) {
        if (browser) {
            localStorage.setItem('auth_token', token);
            localStorage.setItem('auth_user', JSON.stringify(user));
        }
        
        authStore.update(state => ({
            ...state,
            token,
            user,
            isAuthenticated: true,
            error: null
        }));
    },

    // Clear authentication data
    logout() {
        if (browser) {
            localStorage.removeItem('auth_token');
            localStorage.removeItem('auth_user');
        }
        
        authStore.set(initialState);
    },

    // Set loading state
    setLoading(isLoading: boolean) {
        authStore.update(state => ({
            ...state,
            isLoading
        }));
    },

    // Set error
    setError(error: string | null) {
        authStore.update(state => ({
            ...state,
            error,
            isLoading: false
        }));
    },

    // Clear error
    clearError() {
        authStore.update(state => ({
            ...state,
            error: null
        }));
    }
};
