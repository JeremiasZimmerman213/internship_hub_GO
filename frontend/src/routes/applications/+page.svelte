<script lang="ts">
  import { onMount } from "svelte";
  import ApplicationCard from "$lib/components/ApplicationCard.svelte";
  import { apiService } from "$lib/services/apiService";
  import { goto } from "$app/navigation";

  let applications: any[] = [];
  let loading = true;
  let error = "";
  let deleteLoading = false;
  let updatingStatuses: Set<number> = new Set(); // Track which applications are being updated

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

  async function loadApplications() {
    try {
      loading = true;
      error = "";
      applications = await apiService.getApplications();
    } catch (err) {
      error =
        err instanceof Error ? err.message : "Failed to load applications";
      console.error("Error loading applications:", err);
    } finally {
      loading = false;
    }
  }

  async function updateApplicationStatus(applicationId: number, newStatus: number) {
    console.log(`Updating application ${applicationId} to status ${newStatus}`);
    
    try {
      // Add to updating set to show loading state
      updatingStatuses.add(applicationId);
      updatingStatuses = updatingStatuses; // Trigger reactivity

      console.log('Calling API service...');
      // Update the status via API
      const result = await apiService.updateApplicationStatus(applicationId, newStatus);
      console.log('API call successful:', result);
      
      // Update the local state
      applications = applications.map(app => 
        app.id === applicationId 
          ? { ...app, status: newStatus }
          : app
      );
    } catch (err) {
      console.error('Full error object:', err);
      error = err instanceof Error ? err.message : "Failed to update application status";
      console.error("Error updating application status:", err);
    } finally {
      // Remove from updating set
      updatingStatuses.delete(applicationId);
      updatingStatuses = updatingStatuses; // Trigger reactivity
    }
  }

  async function handleDelete(id: number) {
    if (deleteLoading) return;

    try {
      deleteLoading = true;
      await apiService.deleteApplication(id);
      // Remove the deleted application from the list
      applications = applications.filter((app) => app.id !== id);
    } catch (err) {
      error =
        err instanceof Error ? err.message : "Failed to delete application";
      console.error("Error deleting application:", err);
    } finally {
      deleteLoading = false;
    }
  }

  onMount(() => {
    loadApplications();
  });
</script>

<svelte:head>
  <title>Applications - Internship Tracker</title>
</svelte:head>

<div class="container-fluid mt-4">
  <div class="d-flex justify-content-between align-items-center mb-4">
    <div>
      <h1 class="h2 fw-bold" style="color: #4b4b4b;">My Applications</h1>
      <p class="text-muted">Track and manage your internship applications</p>
    </div>
    <div class="d-none d-md-block">
      <a
        href="/applications/new"
        class="btn"
        style="background-color: #B1B2FF; color: #2d2d2d; font-weight: 600;"
      >
        <i class="bi bi-plus-circle"></i> Add New Application
      </a>
    </div>
  </div>

  {#if error}
    <div class="alert alert-danger mb-4">
      <strong>Error:</strong>
      {error}
      <button
        class="btn btn-sm btn-outline-danger ms-2"
        on:click={loadApplications}
      >
        Try Again
      </button>
    </div>
  {/if}

  {#if loading}
    <div class="text-center py-5">
      <div class="spinner-border" style="color: #B1B2FF;" role="status">
        <span class="visually-hidden">Loading...</span>
      </div>
      <p class="mt-2 text-muted">Loading applications...</p>
    </div>
  {:else if applications.length === 0}
    <div class="text-center py-5">
      <div class="mb-4">
        <i class="bi bi-inbox" style="font-size: 4rem; color: #B1B2FF;"></i>
      </div>
      <h4 style="color: #4b4b4b;">No Applications Yet</h4>
      <p class="text-muted mb-4">
        Start tracking your internship applications by adding your first one!
      </p>
      <a
        href="/applications/new"
        class="btn"
        style="background-color: #B1B2FF; color: #2d2d2d; font-weight: 600;"
      >
        <i class="bi bi-plus-circle"></i> Add Your First Application
      </a>
    </div>
  {:else}
    <!-- Desktop Table View -->
    <div class="d-none d-md-block">
      <div class="table-responsive">
        <table class="table table-hover">
          <thead style="background-color: #B1B2FF; color: #2d2d2d;">
            <tr>
              <th>Company</th>
              <th>Position</th>
              <th>Status</th>
              <th>Location</th>
              <th>Applied Date</th>
              <th>Term</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {#each applications as application (application.id)}
              <tr
                style="background-color: #EEF1FF; cursor: pointer;"
                on:click={() => goto(`/applications/${application.id}`)}
                class="clickable-row"
              >
                <td class="fw-semibold" style="color: #4b4b4b;">
                  {application.company}
                </td>
                <td>{application.position}</td>
                <td on:click|stopPropagation>
                  <div class="status-dropdown">
                    {#if updatingStatuses.has(application.id)}
                      <span class="badge bg-secondary">
                        <span class="spinner-border spinner-border-sm me-1" role="status"></span>
                        Updating...
                      </span>
                    {:else}
                      <select
                        class="form-select form-select-sm status-select bg-{statusColors[
                          application.status as keyof typeof statusColors
                        ]}"
                        value={application.status}
                        on:change={(e) => {
                          const target = e.target as HTMLSelectElement | null;
                          if (target) {
                            updateApplicationStatus(application.id, parseInt(target.value));
                          }
                        }}
                      >
                        {#each statusOptions as option}
                          <option value={option.value}>{option.label}</option>
                        {/each}
                      </select>
                    {/if}
                  </div>
                </td>
                <td>{application.location}</td>
                <td>{formatDate(application.applied_date)}</td>
                <td>{application.term}</td>
                <td>
                  <div class="d-flex gap-1">
                    <button
                      class="btn btn-sm btn-outline-primary"
                      title="Edit Application"
                      aria-label="Edit Application"
                      on:click|stopPropagation={() =>
                        goto(`/applications/${application.id}/edit`)}
                    >
                      <i class="bi bi-pencil"></i>
                    </button>
                    {#if application.resume_url}
                      <a
                        href="http://localhost:8080{application.resume_url}"
                        target="_blank"
                        class="btn btn-sm btn-outline-info"
                        title="View Resume"
                        aria-label="View Resume PDF"
                        on:click|stopPropagation
                      >
                        <i class="bi bi-file-pdf"></i>
                      </a>
                    {/if}
                    <button
                      class="btn btn-sm btn-outline-danger"
                      on:click|stopPropagation={() => {
                        if (
                          confirm(
                            `Are you sure you want to delete the application for ${application.position} at ${application.company}? This action cannot be undone.`
                          )
                        ) {
                          handleDelete(application.id);
                        }
                      }}
                      disabled={deleteLoading}
                      title="Delete Application"
                      aria-label="Delete Application"
                    >
                      <i class="bi bi-trash"></i>
                    </button>
                  </div>
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    </div>

    <!-- Mobile Card View -->
    <div class="d-md-none">
      <div class="row g-3">
        {#each applications as application (application.id)}
          <div class="col-12">
            <ApplicationCard {application} onDelete={handleDelete} onStatusUpdate={updateApplicationStatus} />
          </div>
        {/each}
      </div>
    </div>

    <div class="mt-4 text-center text-muted">
      <small>Total Applications: {applications.length}</small>
    </div>
  {/if}
</div>

<!-- Floating Add Button (Mobile) -->
<div class="floating-add-btn d-md-none">
  <a
    href="/applications/new"
    class="btn btn-lg rounded-circle shadow"
    style="background-color: #B1B2FF; color: #2d2d2d;"
    aria-label="Add New Application"
  >
    <i class="bi bi-plus" style="font-size: 1.5rem;"></i>
  </a>
</div>

<style>
  .table {
    border-radius: 1rem;
    overflow: hidden;
    border: none;
  }

  .table th {
    border: none;
    font-weight: 600;
    letter-spacing: 0.5px;
  }

  .table td {
    border: none;
    border-bottom: 1px solid rgba(177, 178, 255, 0.2);
    vertical-align: middle;
  }

  .table tr:last-child td {
    border-bottom: none;
  }

  .floating-add-btn {
    position: fixed;
    bottom: 2rem;
    right: 2rem;
    z-index: 1000;
  }

  .floating-add-btn .btn {
    width: 60px;
    height: 60px;
    display: flex;
    align-items: center;
    justify-content: center;
    transition:
      transform 0.2s,
      box-shadow 0.2s;
  }

  .floating-add-btn .btn:hover {
    transform: scale(1.1);
    background-color: #aac4ff !important;
    box-shadow: 0 8px 20px rgba(177, 178, 255, 0.4) !important;
  }

  .alert {
    border-radius: 0.75rem;
    border: none;
  }

  .btn:hover {
    background-color: #aac4ff !important;
  }

  .clickable-row:hover {
    background-color: #d6dbff !important;
    transform: translateY(-1px);
    box-shadow: 0 2px 8px rgba(177, 178, 255, 0.2);
  }

  .clickable-row {
    transition: all 0.2s ease;
  }

  /* Status Dropdown Styles */
  .status-dropdown {
    min-width: 140px;
  }

  .status-select {
    border: none;
    font-weight: 600;
    font-size: 0.875rem;
    color: white !important;
    cursor: pointer;
    border-radius: 0.375rem;
    padding: 0.375rem 0.75rem;
  }

  .status-select:focus {
    box-shadow: 0 0 0 0.2rem rgba(177, 178, 255, 0.25);
    border-color: #B1B2FF;
  }

  .status-select option {
    background-color: white;
    color: #495057;
  }

  /* Status-specific colors for selects */
  .status-select.bg-primary {
    background-color: #0d6efd !important;
  }

  .status-select.bg-info {
    background-color: #0dcaf0 !important;
  }

  .status-select.bg-warning {
    background-color: #ffc107 !important;
    color: #212529 !important;
  }

  .status-select.bg-success {
    background-color: #198754 !important;
  }

  .status-select.bg-danger {
    background-color: #dc3545 !important;
  }
</style>
