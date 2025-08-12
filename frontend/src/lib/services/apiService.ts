import type { Application, ApiResponse, LoadingState } from '$lib/types/application';
import { writable, get } from 'svelte/store';
import { authStore } from '$lib/stores/authStore';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

// Global loading state store
export const globalLoading = writable<LoadingState>({
    isLoading: false,
    error: null
});

class ApiService {
    private baseURL: string;

    constructor(baseURL: string) {
        this.baseURL = baseURL;
    }

    // Generic request method with error handling and loading states
    private async request<T>(
        endpoint: string,
        options: RequestInit = {},
        showGlobalLoading = false
    ): Promise<T> {
        const url = `${this.baseURL}${endpoint}`;

        if (showGlobalLoading) {
            globalLoading.update(state => ({ ...state, isLoading: true, error: null }));
        }

        try {
            // Get auth token from store
            const authState = get(authStore);
            const authHeaders: Record<string, string> = {};
            
            if (authState.token) {
                authHeaders['Authorization'] = `Bearer ${authState.token}`;
            }

            const response = await fetch(url, {
                ...options,
                headers: {
                    ...authHeaders,
                    ...options.headers,
                },
            });

            if (!response.ok) {
                const errorData = await response.json().catch(() => ({
                    error: `HTTP ${response.status}: ${response.statusText}`
                }));
                throw new Error(errorData.error || `HTTP ${response.status}: ${response.statusText}`);
            }

            // Handle empty responses (like DELETE)
            const contentType = response.headers.get('content-type');
            if (contentType && contentType.includes('application/json')) {
                return await response.json();
            }

            return { success: true } as T;
        } catch (error) {
            const errorMessage = error instanceof Error ? error.message : 'Unknown error occurred';

            if (showGlobalLoading) {
                globalLoading.update(state => ({ isLoading: false, error: errorMessage }));
            }

            throw error;
        } finally {
            if (showGlobalLoading) {
                globalLoading.update(state => ({ ...state, isLoading: false }));
            }
        }
    }

    // Application-specific methods
    async getApplications(): Promise<Application[]> {
        return this.request<Application[]>('/applications', {}, true);
    }

    async getApplication(id: string | number): Promise<Application> {
        return this.request<Application>(`/applications/${id}`);
    }

    async createApplication(formData: FormData): Promise<Application> {
        return this.request<Application>('/applications', {
            method: 'POST',
            body: formData,
        });
    }

    async updateApplication(id: string | number, formData: FormData): Promise<Application> {
        return this.request<Application>(`/applications/${id}`, {
            method: 'PUT',
            body: formData,
        });
    }

    async deleteApplication(id: string | number): Promise<{ success: boolean }> {
        return this.request<{ success: boolean }>(`/applications/${id}`, {
            method: 'DELETE',
        });
    }

    // Utility methods
    clearGlobalError() {
        globalLoading.update(state => ({ ...state, error: null }));
    }

    // File download helper
    getFileUrl(resumeUrl: string): string {
        return `${this.baseURL}${resumeUrl}`;
    }
}

// Export singleton instance
export const apiService = new ApiService(API_BASE_URL);

// Export individual functions for backward compatibility
export const {
    getApplications,
    getApplication,
    createApplication,
    updateApplication,
    deleteApplication,
} = apiService;
