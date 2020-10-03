<script>
    import {dateStore} from "../../store";

    function formatter(d) {
        return `${d.getFullYear()}年${d.getMonth()+1}月`
    }

    function toPrev() {
        dateStore.update(d => new Date(d.setMonth(d.getMonth()-1)))
    }

    function toNext() {
        dateStore.update(d => new Date(d.setMonth(d.getMonth()+1)))
    }
</script>

<div class="controller">
    <p class="column arrow" on:click={toPrev}><span>&lt;</span></p>
    <p class="column date">{formatter($dateStore)}</p>
    <p class="column arrow" on:click={toNext}><span>&gt;</span></p>
</div>

<style>
    .column {
        display: inline-block;
    }
    .date {
        margin: 0 1em;
    }
    .arrow span {
        background-color: #f0f0f0;
        border-radius: 50%;
        width: 1.5em;
        height: 1.5em;
        display: block;
        text-align: center;
    }
    .arrow span:hover {
        background-color: #d0d0d0;
        cursor: pointer;
    }
</style>
