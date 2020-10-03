<style>
    .dropdown {
        position: relative;
        display: inline-block;
    }

    .dropdown-content {
        display: none;
        position: absolute;
        background-color: #f9f9f9;
        width: 150px;
        max-width: 150px;
        box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2);
        padding: .5em;
        z-index: 1;
        overflow: scroll;
        max-height: fit-content;
        white-space: nowrap;
        font-size: .9em;
    }

    .dropdown:hover .dropdown-content {
        display: block;
    }
</style>

<div class="dropdown">
    <slot name="trigger" />
    <div class="dropdown-content">
        <slot />
    </div>
</div>