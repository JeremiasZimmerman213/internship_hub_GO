
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

<svelte:head>
  <title>{isLogin ? 'Login' : 'Register'} - Internship Hub</title>
</svelte:head>

<div class="auth-page">
  <!-- Back Button -->
  <a href="/" class="back-button">
    <i class="bi bi-arrow-left me-2"></i>
    Back to Home
  </a>
  
  <div class="container-fluid d-flex align-items-center justify-content-center min-vh-100 p-4">
    <div class="auth-container">
      <div class="row g-0 h-100">
        <!-- Left Side - Branding/Info -->
        <div class="col-lg-6 d-none d-lg-flex auth-info-section align-items-center justify-content-center">
          <div class="d-flex flex-column text-white p-5">
            <div class="text-center mb-4">
              <i class="bi bi-briefcase-fill mb-3" style="font-size: 4rem;"></i>
              <h1 class="display-5 fw-bold mb-3">Internship Hub</h1>
              <p class="lead mb-0">Your career journey starts here</p>
            </div>
            
            <div class="features-list">
              <div class="feature-item mb-3">
                <i class="bi bi-check-circle-fill me-3"></i>
                <span>Track all your applications</span>
              </div>
              <div class="feature-item mb-3">
                <i class="bi bi-check-circle-fill me-3"></i>
                <span>Manage your documents</span>
              </div>
              <div class="feature-item">
                <i class="bi bi-check-circle-fill me-3"></i>
                <span>Monitor your progress</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Right Side - Form -->
        <div class="col-lg-6 d-flex align-items-center justify-content-center">
          <div class="auth-form-container">
            <div class="form-content-wrapper">
              <div class="text-center mb-5">
                <h2 class="h1 fw-bold mb-3" style="color: #2d2d2d;">
                  {isLogin ? 'Welcome Back' : 'Create Account'}
                </h2>
                <p class="text-muted lead">
                  {isLogin ? 'Sign in to continue your journey' : 'Join us and start tracking your applications'}
                </p>
              </div>

              <form on:submit|preventDefault={handleSubmit} class="auth-form">
                <div class="form-floating mb-4">
                  <input 
                    type="text" 
                    class="form-control form-control-lg" 
                    id="username" 
                    placeholder="Username"
                    bind:value={username} 
                    required 
                    autocomplete="username" 
                  />
                  <label for="username">Username</label>
                </div>

                <div class="form-floating mb-4">
                  <input 
                    type="password" 
                    class="form-control form-control-lg" 
                    id="password" 
                    placeholder="Password"
                    bind:value={password} 
                    required 
                    autocomplete={isLogin ? 'current-password' : 'new-password'} 
                  />
                  <label for="password">Password</label>
                </div>

                <!-- Fixed height container for the optional confirm password field -->
                <div class="confirm-password-slot mb-4">
                  {#if !isLogin}
                    <div class="form-floating">
                      <input 
                        type="password" 
                        class="form-control form-control-lg" 
                        id="confirmPassword" 
                        placeholder="Confirm Password"
                        bind:value={confirmPassword} 
                        required 
                        autocomplete="new-password" 
                      />
                      <label for="confirmPassword">Confirm Password</label>
                    </div>
                  {/if}
                </div>

                {#if error}
                  <div class="alert alert-danger d-flex align-items-center mb-4" role="alert">
                    <i class="bi bi-exclamation-triangle-fill me-2"></i>
                    <div>{error}</div>
                  </div>
                {/if}

                <button 
                  type="submit" 
                  class="btn btn-primary btn-lg w-100 mb-4" 
                  disabled={isLoading}
                >
                  {#if isLoading}
                    <span class="spinner-border spinner-border-sm me-2" role="status" aria-hidden="true"></span>
                  {/if}
                  {isLogin ? 'Sign In' : 'Create Account'}
                </button>
              </form>

              <div class="text-center">
                <p class="text-muted mb-2">
                  {isLogin ? "Don't have an account?" : 'Already have an account?'}
                </p>
                <button 
                  class="btn btn-outline-primary btn-lg" 
                  type="button" 
                  on:click={toggleMode}
                >
                  {isLogin ? 'Create Account' : 'Sign In'}
                </button>
              </div>
            </div>

            <!-- Mobile branding -->
            <div class="d-lg-none text-center mt-5 pt-4 border-top">
              <i class="bi bi-briefcase-fill mb-2" style="font-size: 2rem; color: #B1B2FF;"></i>
              <p class="text-muted small">Internship Hub - Your career journey starts here</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

<style>
  .auth-page {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    overflow-y: auto;
  }

  .auth-page::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: url("data:image/svg+xml,%3Csvg width='60' height='60' viewBox='0 0 60 60' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fill-rule='evenodd'%3E%3Cg fill='%23ffffff' fill-opacity='0.03'%3E%3Ccircle cx='30' cy='30' r='4'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E") repeat;
    z-index: 1;
  }

  .back-button {
    position: absolute;
    top: 2rem;
    left: 2rem;
    z-index: 10;
    color: rgba(255, 255, 255, 0.9);
    text-decoration: none;
    font-weight: 500;
    padding: 0.75rem 1.5rem;
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
    border-radius: 2rem;
    border: 1px solid rgba(255, 255, 255, 0.2);
    transition: all 0.3s ease;
    font-size: 0.95rem;
  }

  .back-button:hover {
    color: white;
    background: rgba(255, 255, 255, 0.2);
    transform: translateY(-1px);
    box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
  }

  .auth-container {
    position: relative;
    z-index: 2;
    width: 100%;
    max-width: 1100px;
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(15px);
    border-radius: 2rem;
    box-shadow: 0 25px 50px rgba(0, 0, 0, 0.15);
    overflow: hidden;
    min-height: 600px;
  }

  .auth-info-section {
    background: linear-gradient(135deg, rgba(102, 126, 234, 0.9) 0%, rgba(118, 75, 162, 0.9) 100%);
    position: relative;
  }

  .auth-info-section::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: url("data:image/svg+xml,%3Csvg width='40' height='40' viewBox='0 0 40 40' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fill-rule='evenodd'%3E%3Cg fill='%23ffffff' fill-opacity='0.08'%3E%3Ccircle cx='20' cy='20' r='3'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E") repeat;
    z-index: 1;
  }

  .auth-info-section > * {
    position: relative;
    z-index: 2;
  }

  .feature-item {
    font-size: 1.1rem;
    display: flex;
    align-items: center;
  }

  .feature-item i {
    color: #B1B2FF;
    font-size: 1.2rem;
  }

  .auth-form-container {
    width: 100%;
    padding: 3rem 2.5rem;
  }

  .form-content-wrapper {
    min-height: 520px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }

  .confirm-password-slot {
    min-height: 82px; /* Fixed height to accommodate form-floating input */
    display: flex;
    align-items: flex-start;
  }

  .auth-form {
    margin-bottom: 2rem;
  }

  .form-control {
    border: 2px solid #e9ecef;
    border-radius: 0.75rem;
    padding: 1rem 1.25rem;
    font-size: 1rem;
    transition: all 0.3s ease;
    background: rgba(255, 255, 255, 0.9);
  }

  .form-control:focus {
    border-color: #B1B2FF;
    box-shadow: 0 0 0 0.2rem rgba(177, 178, 255, 0.25);
    background: rgba(255, 255, 255, 1);
  }

  .form-floating > label {
    color: #6c757d;
    font-weight: 500;
  }

  .btn-primary {
    background: linear-gradient(135deg, #B1B2FF 0%, #8b5cf6 100%);
    border: none;
    border-radius: 0.75rem;
    padding: 1rem;
    font-weight: 600;
    font-size: 1.1rem;
    transition: all 0.3s ease;
    box-shadow: 0 4px 15px rgba(177, 178, 255, 0.3);
  }

  .btn-primary:hover {
    transform: translateY(-2px);
    box-shadow: 0 8px 25px rgba(177, 178, 255, 0.4);
    background: linear-gradient(135deg, #a0a1ff 0%, #7c3aed 100%);
  }

  .btn-primary:disabled {
    transform: none;
    opacity: 0.7;
  }

  .btn-outline-primary {
    border: 2px solid #B1B2FF;
    color: #B1B2FF;
    border-radius: 0.75rem;
    padding: 0.75rem 2rem;
    font-weight: 600;
    transition: all 0.3s ease;
    background: rgba(255, 255, 255, 0.7);
  }

  .btn-outline-primary:hover {
    background: #B1B2FF;
    border-color: #B1B2FF;
    color: white;
    transform: translateY(-1px);
  }

  .alert {
    border: none;
    border-radius: 0.75rem;
    background: rgba(220, 53, 69, 0.1);
    color: #dc3545;
    border-left: 4px solid #dc3545;
  }

  @media (max-width: 991px) {
    .auth-container {
      margin: 1rem;
      border-radius: 1.5rem;
      min-height: auto;
    }
    
    .auth-form-container {
      padding: 2rem 1.5rem;
    }
  }

  @media (max-width: 575px) {
    .auth-container {
      margin: 0.5rem;
      border-radius: 1rem;
    }
    
    .auth-form-container {
      padding: 1.5rem 1.25rem;
    }
  }
</style>
