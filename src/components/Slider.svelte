<script>
    export let collapse;
</script>

<style>
    .outer {
        background-color: #e2e2e2;
        display: none;
    }
    .outer.show {
        display: block;
    }
    .scrollable {
        width: 100%;
        overflow-x: scroll;
    }
</style>

<div class="outer" class:show={!collapse}>
    <div class="scrollable">
        <slot />
    </div>
</div>