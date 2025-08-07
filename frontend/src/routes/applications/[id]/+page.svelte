<script lang="ts">
    import { page } from "$app/stores";
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";
    import { getApplication, deleteApplication } from "$lib/services/apiService";

    let application: any = null;
    let loading = true;
    let error = "";
    let deleteLoading = false;

    const statusLabels = {
        0: "Applied",
        1: "Interview",
        2: "Offer Received",
        3: "Rejected",
    };

    const statusColors = {
        0: "primary",
        1: "warning",
        2: "success",
        3: "danger",
    };

    function formatDate(dateString: string) {
        return new Date(dateString).toLocaleDateString("en-US", {
            year: "numeric",
            month: "long",
            day: "numeric",
        });
    }

    async function loadApplication() {
        try {
            loading = true;
            error = "";
            const id = $page.params.id;
            if (!id) {
                error = "Application ID is required";
                return;
            }
            application = await getApplication(id);
        } catch (err) {
            error =
                err instanceof Error
                    ? err.message
                    : "Failed to load application";
            console.error("Error loading application:", err);
        } finally {
            loading = false;
        }
    }

    async function handleDelete() {
        if (deleteLoading || !application) return;

        const confirmed = confirm(
            `Are you sure you want to delete the application for ${application.position} at ${application.company}? This action cannot be undone.`,
        );

        if (!confirmed) return;

        try {
            deleteLoading = true;
            await deleteApplication(application.id);
            goto("/applications");
        } catch (err) {
            error =
                err instanceof Error
                    ? err.message
                    : "Failed to delete application";
            console.error("Error deleting application:", err);
        } finally {
            deleteLoading = false;
        }
    }

    onMount(() => {
        loadApplication();
    });
</script>

<svelte:head>
    <title>
        {application
            ? `${application.position} at ${application.company}`
            : "Application Details"} - Internship Tracker
    </title>
</svelte:head>

<div class="container mt-4">
    <div class="row justify-content-center">
        <div class="col-lg-8">
            <!-- Back Button -->
            <div class="mb-4">
                <a href="/applications" class="btn btn-outline-secondary">
                    ← Back to Applications
                </a>
            </div>

            {#if error}
                <div class="alert alert-danger mb-4">
                    <div class="d-flex align-items-center">
                        <div class="flex-grow-1">
                            <strong>Error:</strong>
                            {error}
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

            {#if loading}
                <div class="text-center py-5">
                    <div
                        class="spinner-border"
                        style="color: #B1B2FF;"
                        role="status"
                    >
                        <span class="visually-hidden">Loading...</span>
                    </div>
                    <p class="mt-2 text-muted">
                        Loading application details...
                    </p>
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
                        The application you're looking for doesn't exist or may
                        have been deleted.
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
                <!-- Application Details Card -->
                <div
                    class="card shadow-sm"
                    style="background-color: #EEF1FF; border: none; border-radius: 1rem;"
                >
                    <div
                        class="card-header"
                        style="background-color: #B1B2FF; color: #2d2d2d; border-radius: 1rem 1rem 0 0;"
                    >
                        <div
                            class="d-flex justify-content-between align-items-center"
                        >
                            <div>
                                <h3 class="mb-1 fw-bold">
                                    {application.position}
                                </h3>
                                <h5 class="mb-0">{application.company}</h5>
                            </div>
                            <span
                                class="badge bg-{statusColors[
                                    application.status as keyof typeof statusColors
                                ]} fs-6"
                            >
                                {statusLabels[
                                    application.status as keyof typeof statusLabels
                                ]}
                            </span>
                        </div>
                    </div>

                    <div class="card-body p-4">
                        <div class="row g-4">
                            <!-- Location -->
                            <div class="col-md-6">
                                <div class="detail-item">
                                    <i
                                        class="bi bi-geo-alt-fill text-primary me-2"
                                    ></i>
                                    <span class="label">Location:</span>
                                    <span class="value"
                                        >{application.location}</span
                                    >
                                </div>
                            </div>

                            <!-- Applied Date -->
                            <div class="col-md-6">
                                <div class="detail-item">
                                    <i
                                        class="bi bi-calendar-fill text-info me-2"
                                    ></i>
                                    <span class="label">Applied Date:</span>
                                    <span class="value"
                                        >{formatDate(
                                            application.applied_date,
                                        )}</span
                                    >
                                </div>
                            </div>

                            <!-- Term -->
                            <div class="col-md-6">
                                <div class="detail-item">
                                    <i
                                        class="bi bi-clock-fill text-warning me-2"
                                    ></i>
                                    <span class="label">Term:</span>
                                    <span class="value">{application.term}</span
                                    >
                                </div>
                            </div>

                            <!-- Resume -->
                            <div class="col-md-6">
                                <div class="detail-item">
                                    <i
                                        class="bi bi-file-pdf-fill text-danger me-2"
                                    ></i>
                                    <span class="label">Resume:</span>
                                    {#if application.resume_url}
                                        <a
                                            href="http://localhost:8080{application.resume_url}"
                                            target="_blank"
                                            class="btn btn-sm btn-outline-primary ms-2"
                                        >
                                            <i class="bi bi-download me-1"
                                            ></i>Download PDF
                                        </a>
                                    {:else}
                                        <span class="text-muted"
                                            >No resume uploaded</span
                                        >
                                    {/if}
                                </div>
                            </div>

                            <!-- Notes -->
                            {#if application.note}
                                <div class="col-12">
                                    <div class="detail-item">
                                        <i
                                            class="bi bi-journal-text text-secondary me-2"
                                        ></i>
                                        <span class="label">Notes:</span>
                                        <div
                                            class="note-content mt-2 p-3"
                                            style="background-color: #ffffff; border-radius: 0.5rem; border-left: 4px solid #B1B2FF;"
                                        >
                                            {application.note}
                                        </div>
                                    </div>
                                </div>
                            {/if}
                        </div>
                    </div>

                    <!-- Action Buttons -->
                    <div
                        class="card-footer"
                        style="background-color: transparent; border-top: 1px solid rgba(177, 178, 255, 0.2); border-radius: 0 0 1rem 1rem;"
                    >
                        <div class="d-flex gap-2 justify-content-end">
                            <button
                                class="btn btn-outline-secondary d-flex align-items-center justify-content-center"
                                on:click={() =>
                                    goto(
                                        `/applications/${application.id}/edit`,
                                    )}
                            >
                                <i class="bi bi-pencil"></i>Edit
                            </button>
                            <button
                                class="btn btn-outline-danger d-flex align-items-center justify-content-center"
                                on:click={handleDelete}
                                disabled={deleteLoading}
                            >
                                {#if deleteLoading}
                                    <span
                                        class="spinner-border spinner-border-sm me-1"
                                        role="status"
                                    ></span>
                                {:else}
                                    <i class="bi bi-trash"></i>
                                {/if}
                                Delete
                            </button>
                        </div>
                    </div>
                </div>

                <!-- Metadata -->
                <div class="mt-4 text-center">
                    <small class="text-muted">
                        Application ID: {application.id} • Created: {formatDate(
                            application.applied_date,
                        )}
                    </small>
                </div>
            {/if}
        </div>
    </div>
</div>

<style>
    .detail-item {
        display: flex;
        align-items: center;
        flex-wrap: wrap;
        gap: 0.5rem;
    }

    .detail-item .label {
        font-weight: 600;
        color: #4b4b4b;
    }

    .detail-item .value {
        color: #6c757d;
    }

    .note-content {
        font-style: italic;
        color: #495057;
        white-space: pre-wrap;
        word-wrap: break-word;
        width: 100%;
        min-height: 80px;
        display: block;
    }

    .card {
        transition: box-shadow 0.2s ease-in-out;
    }

    .btn:hover {
        transform: translateY(-1px);
        transition: transform 0.2s ease-in-out;
    }

    .alert {
        border-radius: 0.75rem;
        border: none;
    }

    @media (max-width: 768px) {
        .detail-item {
            flex-direction: column;
            align-items: flex-start;
        }

        .card-footer .d-flex {
            flex-direction: column;
        }

        .card-footer .btn {
            min-width: 80px;
        }
    }
</style>
