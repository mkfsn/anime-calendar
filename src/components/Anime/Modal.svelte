<script>
    let show = false;
    const open = () => { show = true }
    const close = () => { show = false }
    let slots = $$props.$$slots;
</script>

<span on:click={open} class="trigger">
    <slot name="trigger">
        <!-- Trigger/Open The Modal -->
        <button>Open Modal</button>
    </slot>
</span>
<div class="modal" class:show={show} on:click|self|stopPropagation={close}>
    <!-- Modal content -->
    <div class="modal-content">
        <span class="close" on:click={close}>
            <svg class="svg-icon" viewBox="0 0 20 20">
                <path d="M10.185,1.417c-4.741,0-8.583,3.842-8.583,8.583c0,4.74,3.842,8.582,8.583,8.582S18.768,14.74,18.768,10C18.768,5.259,14.926,1.417,10.185,1.417 M10.185,17.68c-4.235,0-7.679-3.445-7.679-7.68c0-4.235,3.444-7.679,7.679-7.679S17.864,5.765,17.864,10C17.864,14.234,14.42,17.68,10.185,17.68 M10.824,10l2.842-2.844c0.178-0.176,0.178-0.46,0-0.637c-0.177-0.178-0.461-0.178-0.637,0l-2.844,2.841L7.341,6.52c-0.176-0.178-0.46-0.178-0.637,0c-0.178,0.176-0.178,0.461,0,0.637L9.546,10l-2.841,2.844c-0.178,0.176-0.178,0.461,0,0.637c0.178,0.178,0.459,0.178,0.637,0l2.844-2.841l2.844,2.841c0.178,0.178,0.459,0.178,0.637,0c0.178-0.176,0.178-0.461,0-0.637L10.824,10z"></path>
            </svg>
        </span>
        {#if slots.content}
            <slot name="content" />
        {:else}
            <slot name="header" />
            <slot name="body" />
            <slot name="footer" />
        {/if}
    </div>
</div>

<style>
    .trigger {
        cursor: pointer;
    }
    /* The Modal (background) */
    .modal {
        display: none; /* Hidden by default */
        position: fixed; /* Stay in place */
        z-index: 1; /* Sit on top */
        left: 0;
        top: 0;
        width: 100%; /* Full width */
        height: 100%; /* Full height */
        overflow: auto; /* Enable scroll if needed */
        background-color: rgb(0,0,0); /* Fallback color */
        background-color: rgba(0,0,0,0.4); /* Black w/ opacity */
    }

    .modal.show {
        display: block;
    }

    /* Modal Content/Box */
    .modal-content {
        background-color: #fefefe;
        margin: 7vh auto; /* 15% from the top and centered */
        border: 1px solid #888;
        width: 80%; /* Could be more or less, depending on screen size */
        position: relative;
        padding: 0;
        box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2),0 6px 20px 0 rgba(0,0,0,0.19);
        animation-name: animatetop;
        animation-duration: 0.4s;
        max-width: 960px;
    }

    /* The Close Button */
    .close {
        color: black;
        font-size: 2em;
        position: absolute;
        top: 0.1em;
        right: .1em;
        z-index: 999;
    }
    .close svg {
        width: 1em;
        height: 1em;
    }

    .close:hover,
    .close:active,
    .close:focus {
        text-decoration: none;
        cursor: pointer;
    }

    /* Add Animation */
    @keyframes animatetop {
        from {top: -300px; opacity: 0}
        to {top: 0; opacity: 1}
    }
</style>
