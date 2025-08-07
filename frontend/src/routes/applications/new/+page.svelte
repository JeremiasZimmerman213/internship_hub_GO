<script lang="ts">
  import { goto } from '$app/navigation';
  import ApplicationForm from '$lib/components/ApplicationForm.svelte';
  import { apiService } from '$lib/services/apiService';
  
  let isLoading = false;
  let successMessage = '';
  let errorMessage = '';

  async function handleSubmit(formData: FormData) {
    isLoading = true;
    errorMessage = '';
    successMessage = '';
    
    try {
      const result = await apiService.createApplication(formData);
      successMessage = 'Application created successfully!';
      
      // Redirect to applications page after a short delay
      setTimeout(() => {
        goto('/applications');
      }, 1500);
      
    } catch (error) {
      errorMessage = error instanceof Error ? error.message : 'Failed to create application';
    } finally {
      isLoading = false;
    }
  }
</script>

<svelte:head>
  <title>Add New Application - Internship Tracker</title>
</svelte:head>

<div class="container mt-4">
  <div class="row justify-content-center">
    <div class="col-lg-8">
      <div class="mb-4">
        <h1 class="h2 fw-bold" style="color: #4b4b4b;">Add New Application</h1>
        <p class="text-muted">Fill out the form below to track a new internship application.</p>
      </div>

      {#if successMessage}
        <div class="alert alert-success d-flex align-items-center mb-4">
          <div>
            <strong>Success!</strong> {successMessage} Redirecting to applications page...
          </div>
        </div>
      {/if}

      {#if errorMessage}
        <div class="alert alert-danger mb-4">
          <strong>Error:</strong> {errorMessage}
        </div>
      {/if}

      <ApplicationForm onSubmit={handleSubmit} {isLoading} />

      <div class="mt-4 text-center">
        <a href="/applications" class="btn btn-outline-secondary">
          ← Back to Applications
        </a>
      </div>
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
