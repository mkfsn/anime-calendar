<script>
    import Dropdown from "./Dropdown.svelte";
    import {configStore, timetablesStore} from "../../stores"

    export let anime;
    let selection = $configStore.programs[anime.Id]
        ? $configStore.programs[anime.Id].map(i => anime.Channels[i])
        : [];
    let enabled = $configStore.programs[anime.Id]
        ? $configStore.programs[anime.Id].reduce((acc, cur) => ({...acc, [cur]: true}), {})
        : {};

    $configStore.programs[anime.Id] && $configStore.programs[anime.Id].forEach(i => {
        timetablesStore.addProgram(anime, i);
    })

    function handleChange(channelIndex) {
        enabled[channelIndex] = !enabled[channelIndex];
        if (enabled[channelIndex]) {
            timetablesStore.addProgram(anime, channelIndex);
            configStore.addProgram(anime.Id, channelIndex);
        } else {
            timetablesStore.deleteProgram(anime, channelIndex);
            configStore.deleteProgram(anime.Id, channelIndex);
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
        {#each anime.Channels as channel, index (index)}
            <div>
                <label>
                    <input type="checkbox" bind:group={selection} value={channel} on:change={handleChange(index)} /> {channel}
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