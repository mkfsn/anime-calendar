<script>
    import calendarize from "calendarize";

    export let init = new Date();
    export let year = init.getFullYear();
    export let month = init.getMonth();
    export let today = init;

    let prev = calendarize(new Date(year, month-1)),
        curr = calendarize(new Date(year, month)),
        next = calendarize(new Date(year, month+1));

    function toPrev() {

    }

    function toNext() {

    }

    function isToday(day) {
        return today.getDate() === day;
    }
</script>

<style>
    .calendar {
        position: relative;
        height: calc(100vh - 4em);
    }
    .container {
        display: flex;
        flex-direction: column;
        position: absolute;
        height: 100%;
        width: 100%;
        font-size: .8em;
    }
    .week {
        display: flex;
        flex-grow: 1;
        font-variant-numeric: tabular-nums;
    }

    .date {
        flex-grow: 1;
        width: calc(100% / 7);
        padding: .5em;
        border-right: 1px solid #cccccc;
        border-bottom: 1px solid #cccccc;
    }
    .date.today {
        font-weight: 900;
        background-color: #efefef;
    }
    .date.other {
        color: rgba(128, 128, 128, .5);
    }
    .date.other > * {
        opacity: .2;
    }
</style>

<div class="calendar">
    <div class="container">
        {#each { length:6 } as week, week_idx (week_idx)}
            {#if curr[week_idx]}
                <div class="week">
                    {#each { length:7 } as day, day_idx (day_idx)}
                        {#if curr[week_idx][day_idx] !== 0}
                            <div class="date" class:today={isToday(curr[week_idx][day_idx])}>
                                {#if curr[week_idx][day_idx] === 1}
                                    { month + 1 } / { curr[week_idx][day_idx] }
                                {:else}
                                    { curr[week_idx][day_idx] }
                                {/if}
                            </div>
                        {:else if (week_idx < 1)}
                            <div class="date other">{ prev[prev.length - 1][day_idx] }</div>
                        {:else}
                            <div class="date other">{ next[0][day_idx] }</div>
                        {/if}
                    {/each}
                </div>
            {/if}
        {/each}
    </div>
</div>