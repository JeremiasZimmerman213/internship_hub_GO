<script lang="ts">
    import { goto } from '$app/navigation';
    
    export let application: any;
    export let onDelete: (id: number) => Promise<void>;
    export let onStatusUpdate: ((id: number, status: number) => Promise<void>) | undefined = undefined;

    const statusLabels = {
        0: "Applied",
        1: "OA Received",
        2: "Interviewing",
        3: "Accepted",
        4: "Rejected",
    };

    const statusColors = {
        0: "primary",
        1: "info",
        2: "warning",
        3: "success",
        4: "danger",
    };

    const statusOptions = [
        { value: 0, label: 'Applied' },
        { value: 1, label: 'OA Received' },
        { value: 2, label: 'Interviewing' },
        { value: 3, label: 'Accepted' },
        { value: 4, label: 'Rejected' }
    ];

    function formatDate(dateString: string) {
        return new Date(dateString).toLocaleDateString("en-US", {
            year: "numeric",
            month: "short",
            day: "numeric",
        });
    }
</script>

<div 
    class="card h-100 clickable-card" 
    style="background-color: #eef1ff; border: none; border-radius: 1rem; cursor: pointer;"
    role="button"
    tabindex="0"
    on:click={() => goto(`/applications/${application.id}`)}
    on:keydown={(e) => {
        if (e.key === 'Enter' || e.key === ' ') {
            e.preventDefault();
            goto(`/applications/${application.id}`);
        }
    }}
>
    <div class="card-body">
        <div class="d-flex justify-content-between align-items-start mb-2">
            <h5 class="card-title mb-1" style="color: #4b4b4b;">
                {application.company}
            </h5>
            {#if onStatusUpdate}
                <div class="status-dropdown-mobile">
                    <select
                        class="form-select form-select-sm status-select bg-{statusColors[
                            application.status as keyof typeof statusColors
                        ]}"
                        value={application.status}
                        on:change={(e) => {
                            const target = e.target as HTMLSelectElement | null;
                            if (onStatusUpdate && target) {
                                onStatusUpdate(application.id, parseInt(target.value));
                            }
                        }}
                        on:click|stopPropagation
                    >
                        {#each statusOptions as option}
                            <option value={option.value}>{option.label}</option>
                        {/each}
                    </select>
                </div>
            {:else}
                <span
                    class="badge bg-{statusColors[
                        application.status as keyof typeof statusColors
                    ]} fs-6"
                >
                    {statusLabels[application.status as keyof typeof statusLabels]}
                </span>
            {/if}
        </div>

        <h6 class="card-subtitle mb-2 text-muted">{application.position}</h6>

        <div class="mb-2">
            <small class="text-muted">
                <i class="bi bi-geo-alt"></i> {application.location}
            </small>
        </div>

        <div class="mb-2">
            <small class="text-muted">
                <i class="bi bi-calendar"></i> Applied: {formatDate(application.applied_date)}
            </small>
        </div>

        <div class="mb-3">
            <small class="text-muted">
                <i class="bi bi-clock"></i> {application.term}
            </small>
        </div>

        <div class="d-flex gap-2">
            {#if application.resume_url}
                <a
                    href="http://localhost:8080{application.resume_url}"
                    target="_blank"
                    class="btn btn-sm btn-outline-primary"
                    title="View Resume"
                    on:click|stopPropagation
                >
                    <i class="bi bi-file-pdf"></i> Resume
                </a>
            {/if}
            <button
                class="btn btn-sm btn-outline-danger"
                on:click|stopPropagation={() => {
                    if (confirm(`Are you sure you want to delete the application for ${application.position} at ${application.company}? This action cannot be undone.`)) {
                        onDelete(application.id);
                    }
                }}
                title="Delete Application"
                aria-label="Delete Application"
            >
                <i class="bi bi-trash"></i>
            </button>
        </div>
    </div>
</div>

<style>
    .clickable-card:hover {
        background-color: #D6DBFF !important;
    }

    .card {
        transition:
            transform 0.2s,
            box-shadow 0.2s;
    }

    .card:hover {
        transform: translateY(-2px);
        box-shadow: 0 4px 12px rgba(177, 178, 255, 0.3) !important;
    }

    .card-title {
        font-weight: 600;
    }

    .badge {
        font-size: 0.75rem;
        font-weight: 500;
    }

    .btn-sm {
        font-size: 0.8rem;
    }

    /* Mobile Status Dropdown */
    .status-dropdown-mobile {
        min-width: 130px;
    }

    .status-dropdown-mobile .status-select {
        border: none;
        font-weight: 600;
        font-size: 0.75rem;
        color: white !important;
        cursor: pointer;
        border-radius: 0.375rem;
        padding: 0.25rem 0.5rem;
        min-width: 120px;
    }

    .status-dropdown-mobile .status-select:focus {
        box-shadow: 0 0 0 0.2rem rgba(177, 178, 255, 0.25);
    }

    .status-dropdown-mobile .status-select option {
        background-color: white;
        color: #495057;
    }
</style>
