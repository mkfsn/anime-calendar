<script>
    import Modal from "../Anime/Modal.svelte";
    export let program;

    function getRandomColor(ref, min=128, max=248) {
        const modulo = max - min;

        const colors = [0, 0, 0]; // R, G, B
        for (let i = 0; i < ref.length; i++) {
            const old = colors[i % colors.length];
            colors[i % colors.length] += (old + ref.charCodeAt(i)) % modulo;
        }

        let color = '#';
        for (let i = 0; i < colors.length; i++) {
            color += ('00' + (colors[i] % modulo + min).toString(16)).substr(-2);
        }
        return color;
    }
</script>

<Modal>
    <p class="program" style="--background-color: {getRandomColor(program.program)}" slot="trigger">
        <span>[{program.channel}]</span>
        {program.program}
    </p>
</Modal>


<style>
    .program {
        font-size: .8em;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        margin: 0 0 .4em;
        padding: .2em .4em;
        background-color: var(--background-color);
        border-radius: 5px;
    }
</style>