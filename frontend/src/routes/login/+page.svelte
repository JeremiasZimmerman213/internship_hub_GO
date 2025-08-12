
<script lang="ts">
  import { goto } from '$app/navigation';
  import { authService } from '$lib/services/authService';
  import { authStore } from '$lib/stores/authStore';
  
  let isLogin = true;
  let username = '';
  let password = '';
  let confirmPassword = '';
  let error = '';
  let isLoading = false;

  function toggleMode() {
    isLogin = !isLogin;
    error = '';
    username = '';
    password = '';
    confirmPassword = '';
  }

  async function handleSubmit(event: Event) {
    event.preventDefault();
    error = '';
    
    if (!username || !password) {
      error = 'Please fill in all required fields.';
      return;
    }
    
    if (!isLogin && password !== confirmPassword) {
      error = 'Passwords do not match.';
      return;
    }

    isLoading = true;
    
    try {
      if (isLogin) {
        // Login
        await authService.login({ username, password });
        goto('/applications'); // Redirect to applications page
      } else {
        // Register
        await authService.register({ username, password });
        alert('Registration successful! Please log in.');
        isLogin = true; // Switch to login mode
        password = '';
        confirmPassword = '';
      }
    } catch (err) {
      error = err instanceof Error ? err.message : 'An error occurred';
    } finally {
      isLoading = false;
    }
  }
</script>

<div class="auth-container">
  <div class="card shadow-sm p-4">
    <h2 class="mb-3 text-center fw-bold" style="color: #4b4b4b;">
      {isLogin ? 'Login' : 'Register'}
    </h2>
    <form on:submit|preventDefault={handleSubmit}>
      <div class="mb-3">
        <label for="username" class="form-label">Username</label>
        <input type="text" class="form-control" id="username" bind:value={username} required autocomplete="username" />
      </div>
      <div class="mb-3">
        <label for="password" class="form-label">Password</label>
        <input type="password" class="form-control" id="password" bind:value={password} required autocomplete={isLogin ? 'current-password' : 'new-password'} />
      </div>
      {#if !isLogin}
        <div class="mb-3">
          <label for="confirmPassword" class="form-label">Confirm Password</label>
          <input type="password" class="form-control" id="confirmPassword" bind:value={confirmPassword} required autocomplete="new-password" />
        </div>
      {/if}
      {#if error}
        <div class="alert alert-danger py-2">{error}</div>
      {/if}
      <button type="submit" class="btn w-100 mb-2" style="background-color: #B1B2FF; color: #2d2d2d; font-weight: 600;" disabled={isLoading}>
        {#if isLoading}
          <span class="spinner-border spinner-border-sm me-2" role="status"></span>
        {/if}
        {isLogin ? 'Login' : 'Register'}
      </button>
    </form>
    <div class="text-center mt-2">
      <button class="btn btn-link p-0" type="button" on:click={toggleMode}>
        {isLogin ? "Don't have an account? Register" : 'Already have an account? Login'}
      </button>
    </div>
  </div>
</div>

<style>
.auth-container {
  max-width: 400px;
  margin: 3rem auto 0 auto;
}
.card {
  border: none;
  background: #EEF1FF;
  border-radius: 1rem;
}
.btn-link {
  color: #6c63ff;
  text-decoration: underline;
}
.btn-link:hover {
  color: #B1B2FF;
}
</style>
