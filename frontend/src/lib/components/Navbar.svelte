<script lang="ts">
  import { authStore, authActions } from '$lib/stores/authStore';
  import { goto } from '$app/navigation';

  function handleLogout() {
    authActions.logout();
    goto('/login');
  }
</script>

<nav class="navbar navbar-expand-lg navbar-light" style="background-color: #B1B2FF;">
  <div class="container-fluid">
    <a class="navbar-brand fw-bold" href="/" style="color: #2d2d2d;">Internship Hub</a>
    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarNav">
      <ul class="navbar-nav ms-auto">
        {#if $authStore.isAuthenticated}
          <!-- Authenticated user menu -->
          <li class="nav-item">
            <a class="nav-link" href="/applications" style="color: #2d2d2d;">Applications</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="/applications/new" style="color: #2d2d2d;">Add New</a>
          </li>
          <li class="nav-item dropdown">
            <!-- svelte-ignore a11y_invalid_attribute -->
            <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false" style="color: #2d2d2d;">
              Welcome, {$authStore.user?.username}
            </a>
            <ul class="dropdown-menu">
              <li><button class="dropdown-item" on:click={handleLogout}>Logout</button></li>
            </ul>
          </li>
        {:else}
          <!-- Non-authenticated user menu -->
          <li class="nav-item">
            <a class="nav-link" href="/login" style="color: #2d2d2d;">Login</a>
          </li>
        {/if}
      </ul>
    </div>
  </div>
</nav>

<style>
nav.navbar {
  /* Use palette for gradient or accents if desired */
  box-shadow: 0 2px 8px rgba(0,0,0,0.04);
}
.navbar-brand {
  letter-spacing: 1px;
}
.nav-link {
  margin-left: 0.5rem;
  margin-right: 0.5rem;
  transition: background 0.2s, color 0.2s;
  border-radius: 4px;
}
.nav-link:hover {
  background-color: #AAC4FF;
  color: #2d2d2d;
}
.dropdown-item {
  border: none;
  background: none;
  width: 100%;
  text-align: left;
}
.dropdown-item:hover {
  background-color: #f8f9fa;
}
</style>
