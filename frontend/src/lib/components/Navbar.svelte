<script>
  // No props or state needed for static header
</script>

<nav class="navbar navbar-expand-lg navbar-light" style="background-color: #B1B2FF;">
  <div class="container-fluid">
    <a class="navbar-brand fw-bold" href="/" style="color: #2d2d2d;">Internship Hub</a>
    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarNav">
      <ul class="navbar-nav ms-auto">
        <li class="nav-item">
          <a class="nav-link" href="/applications" style="color: #2d2d2d;">Applications</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="/applications/new" style="color: #2d2d2d;">Add New</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="/login" style="color: #2d2d2d;">Login</a>
        </li>
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
.nav-link:hover, .nav-link.active {
  background-color: #AAC4FF;
  color: #2d2d2d;
}
</style>
