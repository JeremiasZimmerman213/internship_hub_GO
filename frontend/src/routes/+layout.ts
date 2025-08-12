import { redirect } from '@sveltejs/kit';

export const load = async ({ url, fetch }: { url: URL; fetch: any }) => {
    // List of protected routes
    const protectedRoutes = ['/applications'];
    
    // Check if current route is protected
    const isProtectedRoute = protectedRoutes.some(route => 
        url.pathname.startsWith(route)
    );

    if (isProtectedRoute) {
        // Check for auth token in localStorage (client-side only)
        if (typeof window !== 'undefined') {
            const token = localStorage.getItem('auth_token');
            
            if (!token) {
                // No token found, redirect to login
                throw redirect(302, '/login');
            }

            // Optional: Verify token with backend
            try {
                const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';
                const response = await fetch(`${API_BASE_URL}/user/profile`, {
                    headers: {
                        'Authorization': `Bearer ${token}`
                    }
                });

                if (!response.ok) {
                    // Token is invalid, clear it and redirect
                    localStorage.removeItem('auth_token');
                    localStorage.removeItem('auth_user');
                    throw redirect(302, '/login');
                }
            } catch (error) {
                // Network error or invalid token
                localStorage.removeItem('auth_token');
                localStorage.removeItem('auth_user');
                throw redirect(302, '/login');
            }
        }
    }

    return {};
};
