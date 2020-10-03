<script lang="ts">
    import {timetables} from "../store";

    export let month = -1;
    export let day = -1;
    export let isCurr = false;

    let isToday = (new Date).getDate() === day;
    let programs = [];

    timetables.subscribe(_programs => {
        programs = Object.keys(_programs).reduce((prev, program) => {
            Object.entries(_programs[program]).forEach(([channel, episodes]) => {
                episodes.forEach(episode => {
                    const d = new Date(episode.StartAt);
                    if (d.getMonth() === month && d.getDate() === day) {
                        prev.push({channel, program});
                    }
                })
            })
            return prev;
        }, []);
    });
</script>

<div class="date" class:today={isToday} class:other={!isCurr}>
    {#if day === 1}
        <p class="title">{ month + 1 } / { day }</p>
    {:else}
        <p class="title">{ day }</p>
    {/if}
    <div class="programs">
        {#each programs as program}
            <p class="program">
                <span>[{program.channel}]</span>
                {program.program}
            </p>
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
    .date.other {
        color: rgba(126, 126, 126, 0.5);
    }
    .date.other > * {
        opacity: .2;
    }
    .date.other > .title {
        opacity: 1;
    }
    .title {
        margin-top: 0;
        margin-bottom: .5em;
    }
    .program {
        font-size: .8em;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        margin: 0 0 .4em;
        padding: .2em .4em;
        background-color: aquamarine;
        border-radius: 5px;
    }
</style>