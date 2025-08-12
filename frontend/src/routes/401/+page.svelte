<script lang="ts">
  import { goto } from '$app/navigation';
</script>

<svelte:head>
  <title>Unauthorized - Internship Hub</title>
</svelte:head>

<div class="container mt-5">
  <div class="row justify-content-center">
    <div class="col-md-6 text-center">
      <div class="mb-4">
        <i class="bi bi-shield-exclamation" style="font-size: 4rem; color: #dc3545;"></i>
      </div>
      
      <h1 class="h2 fw-bold mb-3" style="color: #4b4b4b;">401 - Unauthorized</h1>
      
      <p class="lead text-muted mb-4">
        You need to be logged in to access this page.
      </p>
      
      <p class="text-muted mb-4">
        Please log in with your account to continue using the Internship Hub.
      </p>
      
      <div class="d-flex gap-2 justify-content-center">
        <button 
          on:click={() => goto('/login')}
          class="btn"
          style="background-color: #B1B2FF; color: #2d2d2d; font-weight: 600;"
        >
          <i class="bi bi-box-arrow-in-right me-2"></i>
          Go to Login
        </button>
        
        <button 
          on:click={() => goto('/')}
          class="btn btn-outline-secondary"
        >
          <i class="bi bi-house me-2"></i>
          Go Home
        </button>
      </div>
    </div>
  </div>
</div>

<style>
  .container {
    max-width: 600px;
  }
  
  .btn:hover {
    transform: translateY(-1px);
    transition: transform 0.2s ease-in-out;
  }
</style>
