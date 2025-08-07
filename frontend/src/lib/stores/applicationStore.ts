import { writable, derived } from 'svelte/store';
import type { Application, LoadingState } from '$lib/types/application';
import { apiService } from '$lib/services/apiService';

// Application data store
export const applications = writable<Application[]>([]);

// Loading state for applications
export const applicationsLoading = writable<LoadingState>({
    isLoading: false,
    error: null
});

// Derived stores for filtered/sorted data
export const applicationsByStatus = derived(applications, ($applications) => {
    return $applications.reduce((acc, app) => {
        if (!acc[app.status]) {
            acc[app.status] = [];
        }
        acc[app.status].push(app);
        return acc;
    }, {} as Record<number, Application[]>);
});

// Application store actions
export const applicationActions = {
    // Load all applications
    async loadApplications() {
        applicationsLoading.update(state => ({ ...state, isLoading: true, error: null }));

        try {
            const data = await apiService.getApplications();
            applications.set(data);
        } catch (error) {
            const errorMessage = error instanceof Error ? error.message : 'Failed to load applications';
            applicationsLoading.update(state => ({ isLoading: false, error: errorMessage }));
            throw error;
        } finally {
            applicationsLoading.update(state => ({ ...state, isLoading: false }));
        }
    },

    // Add new application
    async addApplication(formData: FormData) {
        try {
            const newApp = await apiService.createApplication(formData);
            applications.update(apps => [...apps, newApp]);
            return newApp;
        } catch (error) {
            throw error;
        }
    },

    // Update existing application
    async updateApplication(id: string | number, formData: FormData) {
        try {
            const updatedApp = await apiService.updateApplication(id, formData);
            applications.update(apps =>
                apps.map(app => app.id === updatedApp.id ? updatedApp : app)
            );
            return updatedApp;
        } catch (error) {
            throw error;
        }
    },

    // Delete application
    async deleteApplication(id: string | number) {
        try {
            await apiService.deleteApplication(id);
            applications.update(apps => apps.filter(app => app.id !== Number(id)));
        } catch (error) {
            throw error;
        }
    },

    // Clear error state
    clearError() {
        applicationsLoading.update(state => ({ ...state, error: null }));
    }
};
