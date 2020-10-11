<script lang="ts">
    import {afterUpdate} from "svelte";
    import {timetables} from "../../store";
    import CalendarProgram from "./CalendarProgram.svelte";

    export let date = new Date(); // Date
    export let day = -1;
    export let isCurr = false;

    let programs = [];
    let today = new Date();
    let isToday = (month, day) => {
        return today.getDate() === day &&
            today.getMonth() === month &&
            today.getFullYear() === date.getFullYear();
    }

    let allPrograms = {};

    function updatePrograms() {
        programs = Object.values(allPrograms).reduce((acc, program) => {
            program.selected.forEach((channelIndex) => {
                program.Timetables[channelIndex].forEach(episode => {
                    const d = new Date(episode.StartAt);
                    if (d.getFullYear() === date.getFullYear() && d.getMonth() === date.getMonth() && d.getDate() === day) {
                        acc.push({program, episode});
                    }
                })
            })
            return acc;
        }, []);
    }

    timetables.subscribe(_programs => {
        allPrograms = _programs;
        updatePrograms()
    });
    afterUpdate(() => {
        updatePrograms();
    })
</script>

<div class="date" class:today={isToday(date.getMonth(), day)} class:other={!isCurr}>
    {#if day === 1}
        <p class="title">{ date.getMonth() + 1 } / { day }</p>
    {:else}
        <p class="title">{ day }</p>
    {/if}
    <div class="programs">
        {#each programs as program}
            <CalendarProgram program={program.program} episode={program.episode} />
        {/each}
    </div>
</div>

<style>
    .date {
        padding: .4em .2em;
        border-right: 1px solid #cccccc;
        border-bottom: 1px solid #cccccc;
        /*https://stackoverflow.com/questions/40351223/equal-width-flex-items-regardless-of-content-length*/
        flex: 1 1 0px;
        /*https://css-tricks.com/flexbox-truncated-text/*/
        min-width: 1px;
    }
    .date.today {
        font-weight: 900;
        background-color: #efefef;
    }
    .date.other .title {
        color: rgba(126, 126, 126, 0.5);
    }
    .date.other .programs :global(.program) {
        opacity: .5;
    }
    .date.other > .title {
        opacity: 1;
    }
    .title {
        margin-top: 0;
        margin-bottom: .5em;
    }
</style>