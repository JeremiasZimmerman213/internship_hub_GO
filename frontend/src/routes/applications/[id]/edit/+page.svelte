<script lang="ts">
    import { page } from "$app/stores";
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";
    import ApplicationForm from "$lib/components/ApplicationForm.svelte";
    import { getApplication, updateApplication } from "$lib/services/apiService";

    let application: any = null;
    let isLoading = false;
    let loadingData = true;
    let successMessage = "";
    let errorMessage = "";

    async function loadApplication() {
        try {
            loadingData = true;
            errorMessage = "";
            const id = $page.params.id;
            if (!id) {
                errorMessage = "Application ID is required";
                return;
            }
            application = await getApplication(id);
        } catch (err) {
            errorMessage =
                err instanceof Error
                    ? err.message
                    : "Failed to load application";
            console.error("Error loading application:", err);
        } finally {
            loadingData = false;
        }
    }

    async function handleSubmit(formData: FormData) {
        if (!application) return;

        isLoading = true;
        errorMessage = "";
        successMessage = "";

        try {
            const result = await updateApplication(application.id, formData);
            successMessage = "Application updated successfully!";

            // Redirect to application details page after a short delay
            setTimeout(() => {
                goto(`/applications/${application.id}`);
            }, 1500);
        } catch (error) {
            errorMessage =
                error instanceof Error
                    ? error.message
                    : "Failed to update application";
        } finally {
            isLoading = false;
        }
    }

    onMount(() => {
        loadApplication();
    });
</script>

<svelte:head>
    <title>
        {application
            ? `Edit ${application.position} at ${application.company}`
            : "Edit Application"} - Internship Tracker
    </title>
</svelte:head>

<div class="container mt-4">
    <div class="row justify-content-center">
        <div class="col-lg-8">
            <div class="mb-4">
                <h1 class="h2 fw-bold" style="color: #4b4b4b;">
                    Edit Application
                </h1>
                <p class="text-muted">
                    {#if application}
                        Editing application for {application.position} at {application.company}
                    {:else}
                        Update your internship application details below.
                    {/if}
                </p>
            </div>

            {#if errorMessage}
                <div class="alert alert-danger mb-4">
                    <div class="d-flex align-items-center">
                        <div class="flex-grow-1">
                            <strong>Error:</strong>
                            {errorMessage}
                        </div>
                        <button
                            class="btn btn-sm btn-outline-danger"
                            on:click={loadApplication}
                        >
                            Try Again
                        </button>
                    </div>
                </div>
            {/if}

            {#if successMessage}
                <div class="alert alert-success d-flex align-items-center mb-4">
                    <div>
                        <strong>Success!</strong>
                        {successMessage} Redirecting to application details...
                    </div>
                </div>
            {/if}

            {#if loadingData}
                <div class="text-center py-5">
                    <div
                        class="spinner-border"
                        style="color: #B1B2FF;"
                        role="status"
                    >
                        <span class="visually-hidden">Loading...</span>
                    </div>
                    <p class="mt-2 text-muted">Loading application data...</p>
                </div>
            {:else if !application}
                <div class="text-center py-5">
                    <div class="mb-4">
                        <i
                            class="bi bi-exclamation-triangle"
                            style="font-size: 4rem; color: #B1B2FF;"
                        ></i>
                    </div>
                    <h4 style="color: #4b4b4b;">Application Not Found</h4>
                    <p class="text-muted mb-4">
                        The application you're trying to edit doesn't exist or
                        may have been deleted.
                    </p>
                    <a
                        href="/applications"
                        class="btn"
                        style="background-color: #B1B2FF; color: #2d2d2d; font-weight: 600;"
                    >
                        View All Applications
                    </a>
                </div>
            {:else}
                <ApplicationForm
                    onSubmit={handleSubmit}
                    {isLoading}
                    initialData={application}
                    isEditMode={true}
                />

                <div class="mt-4 d-flex gap-2 justify-content-center">
                    <a
                        href="/applications/{application.id}"
                        class="btn btn-outline-secondary"
                    >
                        ← Back to Details
                    </a>
                    <a href="/applications" class="btn btn-outline-secondary">
                        All Applications
                    </a>
                </div>
            {/if}
        </div>
    </div>
</div>

<style>
    .container {
        max-width: 900px;
    }

    .alert {
        border-radius: 0.75rem;
        border: none;
    }

    .alert-success {
        background-color: #d1e7dd;
        color: #0a3622;
    }

    .alert-danger {
        background-color: #f8d7da;
        color: #58151c;
    }
</style>
