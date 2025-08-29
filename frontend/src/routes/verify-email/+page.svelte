<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/stores';
  import { authService } from '$lib/services/authService';
  import { goto } from '$app/navigation';

  let message = '';
  let error = '';
  let isLoading = true;

  onMount(async () => {
    const url = new URL(window.location.href);
    const token = url.searchParams.get('token');
    if (!token) {
      error = 'Missing verification token.';
      isLoading = false;
      return;
    }

    try {
      const res = await authService.verifyEmail(token);
      message = res.message || 'Email verified successfully!';
    } catch (e) {
      error = e instanceof Error ? e.message : 'Verification failed.';
    } finally {
      isLoading = false;
    }
  });

  function goToLogin() {
    goto('/login');
  }
</script>

<svelte:head>
  <title>Email Verification - Internship Hub</title>
</svelte:head>

<div class="container py-5">
  <div class="row justify-content-center">
    <div class="col-md-8">
      <div class="card shadow-sm">
        <div class="card-body p-4 text-center">
          <h1 class="h3 mb-3">Email Verification</h1>
          {#if isLoading}
            <div class="spinner-border" role="status">
              <span class="visually-hidden">Loading...</span>
            </div>
          {:else if message}
            <div class="alert alert-success" role="alert">{message}</div>
            <button class="btn btn-primary" on:click={goToLogin}>Go to Login</button>
          {:else}
            <div class="alert alert-danger" role="alert">{error}</div>
            <button class="btn btn-outline-secondary" on:click={goToLogin}>Back to Login</button>
          {/if}
        </div>
      </div>
    </div>
  </div>
</div>
