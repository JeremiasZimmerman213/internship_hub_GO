<script lang="ts">
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
    import { authStore } from '$lib/stores/authStore';

    onMount(() => {
        // Redirect authenticated users to applications
        const unsubscribe = authStore.subscribe(state => {
            if (state.isAuthenticated) {
                goto('/applications');
            }
        });

        return unsubscribe;
    });
</script>

<div class="container mt-5">
    <div class="row justify-content-center">
        <div class="col-lg-8 text-center">
            <div class="mb-5">
                <h1 class="display-4 fw-bold mb-3" style="color: #4b4b4b;">
                    Welcome to Internship Hub
                </h1>
                <p class="lead text-muted mb-4">
                    Track and manage your internship applications in one place
                </p>
            </div>

            <div class="row g-4 mb-5">
                <div class="col-md-4">
                    <div class="card h-100 border-0" style="background-color: #EEF1FF;">
                        <div class="card-body p-4">
                            <i class="bi bi-journal-plus" style="font-size: 3rem; color: #B1B2FF;"></i>
                            <h5 class="mt-3 mb-2">Track Applications</h5>
                            <p class="text-muted">Keep track of all your internship applications with detailed status updates</p>
                        </div>
                    </div>
                </div>
                <div class="col-md-4">
                    <div class="card h-100 border-0" style="background-color: #EEF1FF;">
                        <div class="card-body p-4">
                            <i class="bi bi-file-pdf" style="font-size: 3rem; color: #B1B2FF;"></i>
                            <h5 class="mt-3 mb-2">Manage Resumes</h5>
                            <p class="text-muted">Upload and organize your resumes for different applications</p>
                        </div>
                    </div>
                </div>
                <div class="col-md-4">
                    <div class="card h-100 border-0" style="background-color: #EEF1FF;">
                        <div class="card-body p-4">
                            <i class="bi bi-graph-up" style="font-size: 3rem; color: #B1B2FF;"></i>
                            <h5 class="mt-3 mb-2">Monitor Progress</h5>
                            <p class="text-muted">View your application status and track your progress</p>
                        </div>
                    </div>
                </div>
            </div>

            <div class="mb-4">
                <a 
                    href="/login" 
                    class="btn btn-lg me-3"
                    style="background-color: #B1B2FF; color: #2d2d2d; font-weight: 600;"
                >
                    Get Started
                </a>
                <a href="/login" class="btn btn-outline-secondary btn-lg">
                    Learn More
                </a>
            </div>
        </div>
    </div>
</div>

<style>
    .card {
        transition: transform 0.2s ease-in-out;
    }
    
    .card:hover {
        transform: translateY(-5px);
    }
    
    .btn:hover {
        transform: translateY(-1px);
        transition: transform 0.2s ease-in-out;
    }
</style>