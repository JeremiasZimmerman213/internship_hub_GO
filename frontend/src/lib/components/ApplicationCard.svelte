<script lang="ts">
    export let application: any;
    export let onDelete: (id: number) => void;

    const statusColors = {
        0: "bg-primary", // Applied
        1: "bg-warning", // Interview
        2: "bg-success", // Offer Received
        3: "bg-danger", // Rejected
    };

    const statusLabels = {
        0: "Applied",
        1: "Interview",
        2: "Offer Received",
        3: "Rejected",
    };

    function formatDate(dateString: string) {
        return new Date(dateString).toLocaleDateString("en-US", {
            year: "numeric",
            month: "short",
            day: "numeric",
        });
    }

    function handleDelete() {
        if (
            confirm(
                `Are you sure you want to delete the application for ${application.position} at ${application.company}?`,
            )
        ) {
            onDelete(application.id);
        }
    }
</script>

<div class="application-card card shadow-sm h-100">
    <div class="card-body">
        <div class="d-flex justify-content-between align-items-start mb-2">
            <h5 class="card-title mb-1" style="color: #4b4b4b;">
                {application.company}
            </h5>
            <span class="badge {statusColors[application.status as keyof typeof statusColors]} text-white">
                {statusLabels[application.status as keyof typeof statusLabels]}
            </span>
        </div>

        <h6 class="card-subtitle mb-2 text-muted">{application.position}</h6>

        <div class="application-details">
            <div class="detail-item">
                <i class="bi bi-geo-alt"></i>
                <span>{application.location}</span>
            </div>
            <div class="detail-item">
                <i class="bi bi-calendar"></i>
                <span>Applied: {formatDate(application.applied_date)}</span>
            </div>
            <div class="detail-item">
                <i class="bi bi-clock"></i>
                <span>{application.term}</span>
            </div>
        </div>

        {#if application.note}
            <p class="card-text mt-2 text-muted small">
                {application.note.length > 80
                    ? application.note.substring(0, 80) + "..."
                    : application.note}
            </p>
        {/if}
    </div>

    <div class="card-footer bg-transparent border-top-0">
        <div class="d-flex gap-2">
            {#if application.resume_url}
                <a
                    href="http://localhost:8080{application.resume_url}"
                    target="_blank"
                    class="btn btn-sm btn-outline-primary flex-fill"
                >
                    <i class="bi bi-file-pdf"></i> Resume
                </a>
            {/if}
            <button
                class="btn btn-sm btn-outline-danger"
                on:click={handleDelete}
                aria-label="Delete application"
            >
                <i class="bi bi-trash"></i>
            </button>
        </div>
    </div>
</div>

<style>
    .application-card {
        border: none;
        background: #eef1ff;
        border-radius: 1rem;
        transition:
            transform 0.2s,
            box-shadow 0.2s;
    }

    .application-card:hover {
        transform: translateY(-2px);
        box-shadow: 0 4px 12px rgba(177, 178, 255, 0.3) !important;
    }

    .detail-item {
        display: flex;
        align-items: center;
        gap: 0.5rem;
        margin-bottom: 0.25rem;
        font-size: 0.875rem;
        color: #6c757d;
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
</style>
