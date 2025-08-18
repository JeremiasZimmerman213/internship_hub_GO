<script lang="ts">
    import { createEventDispatcher } from 'svelte';

    export let show = false;
    export let applicationPosition = '';
    export let applicationCompany = '';
    export let isLoading = false;

    const dispatch = createEventDispatcher();

    function handleConfirm() {
        dispatch('confirm');
    }

    function handleCancel() {
        dispatch('cancel');
        show = false;
    }

    // Close modal when clicking backdrop
    function handleBackdropClick(event: MouseEvent) {
        if (event.target === event.currentTarget) {
            handleCancel();
        }
    }

    // Handle escape key
    function handleKeydown(event: KeyboardEvent) {
        if (event.key === 'Escape' && show) {
            handleCancel();
        }
    }
</script>

<svelte:window on:keydown={handleKeydown} />

{#if show}
    <!-- Modal Backdrop -->
    <div 
        class="modal fade show" 
        style="display: block;" 
        tabindex="-1" 
        role="dialog" 
        aria-labelledby="deleteModalLabel" 
        aria-hidden="false"
        on:click={handleBackdropClick}
        on:keydown={handleKeydown}
    >
        <div class="modal-dialog modal-dialog-centered" role="document">
            <div class="modal-content" style="border-radius: 1rem; border: none;">
                <!-- Modal Header -->
                <div class="modal-header" style="background-color: #dc3545; color: white; border-radius: 1rem 1rem 0 0;">
                    <h5 class="modal-title d-flex align-items-center" id="deleteModalLabel">
                        <i class="bi bi-exclamation-triangle-fill me-2"></i>
                        Confirm Delete
                    </h5>
                    <button 
                        type="button" 
                        class="btn-close btn-close-white" 
                        aria-label="Close"
                        on:click={handleCancel}
                        disabled={isLoading}
                    ></button>
                </div>

                <!-- Modal Body -->
                <div class="modal-body p-4">
                    <div class="text-center mb-3">
                        <i class="bi bi-trash3-fill text-danger" style="font-size: 3rem;"></i>
                    </div>
                    <p class="text-center mb-3">
                        Are you sure you want to delete this application?
                    </p>
                    <p class="text-muted text-center small mb-0">
                        <i class="bi bi-info-circle me-1"></i>
                        This action cannot be undone.
                    </p>
                </div>

                <!-- Modal Footer -->
                <div class="modal-footer" style="border-top: 1px solid #dee2e6;">
                    <button 
                        type="button" 
                        class="btn btn-outline-secondary"
                        on:click={handleCancel}
                        disabled={isLoading}
                    >
                        <i class="bi bi-x-circle me-1"></i>
                        Cancel
                    </button>
                    <button 
                        type="button" 
                        class="btn btn-danger"
                        on:click={handleConfirm}
                        disabled={isLoading}
                    >
                        {#if isLoading}
                            <span class="spinner-border spinner-border-sm me-2" role="status"></span>
                            Deleting...
                        {:else}
                            <i class="bi bi-trash3 me-1"></i>
                            Delete Application
                        {/if}
                    </button>
                </div>
            </div>
        </div>
    </div>

    <!-- Modal Backdrop -->
    <div class="modal-backdrop fade show"></div>
{/if}

<style>
    .modal {
        z-index: 1055;
    }

    .modal-backdrop {
        z-index: 1050;
        background-color: rgba(0, 0, 0, 0.5);
    }

    .modal-content {
        box-shadow: 0 0.5rem 1rem rgba(0, 0, 0, 0.15);
    }

    .btn:focus {
        box-shadow: 0 0 0 0.2rem rgba(220, 53, 69, 0.25);
    }

    .btn-outline-secondary:hover {
        transform: translateY(-1px);
        transition: transform 0.2s ease;
    }

    .btn-danger:hover {
        transform: translateY(-1px);
        transition: transform 0.2s ease;
    }

    .modal-dialog {
        transition: transform 0.3s ease-out;
    }

    .alert {
        background-color: #f8f9fa;
        border: 1px solid #dee2e6;
        border-radius: 0.5rem;
    }
</style>
