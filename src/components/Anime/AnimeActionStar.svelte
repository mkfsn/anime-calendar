<script>
    import Dropdown from "./Dropdown.svelte";
    import {timetables} from "../store";

    export let anime;
    let selection = [];
    let enabled = {};
    function handleChange(channel) {
        enabled[channel] = !enabled[channel];
        if (enabled[channel]) {
            timetables.addProgram(anime.Title, channel, anime.Timetables[channel]);
        } else {
            timetables.deleteProgram(anime.Title, channel);
        }
    }
</script>

<div class="star">
    <Dropdown>
        <i class="trigger" class:active={selection.length !== 0} slot="trigger">
            <svg class="svg-icon" viewBox="-1 -1 18 18">
                <path d="M3.612 15.443c-.386.198-.824-.149-.746-.592l.83-4.73L.173 6.765c-.329-.314-.158-.888.283-.95l4.898-.696L7.538.792c.197-.39.73-.39.927 0l2.184 4.327 4.898.696c.441.062.612.636.283.95l-3.523 3.356.83 4.73c.078.443-.36.79-.746.592L8 13.187l-4.389 2.256z"></path>
            </svg>
        </i>
        {#each Array.from(Object.entries(anime.Timetables)) as [channel, timetable]}
            <div>
                <label>
                    <input type="checkbox" bind:group={selection} value={channel} on:change={handleChange(channel)} /> {channel}
                </label>
            </div>
        {:else}
            <span class="no-data"><i>No Data</i></span>
        {/each}
    </Dropdown>
</div>

<style>
    .star {
        text-align: left;
        display: inline-block;
    }
    .trigger {
        width: 1em;
        height: 1em;
        border-radius: 50%;
        padding: .1em;
    }
    .trigger svg path {
        fill: #c8c8c8;
        stroke: #8c8c8c;
    }
    .trigger.active svg path {
        fill: #ffff58;
        stroke: #d0c61a;
    }
    svg {
        width: 1em;
        height: 1em;
        line-height: 1em;
        vertical-align: text-top;
    }
    .no-data {
        color: #9f9f9f;
    }
</style>