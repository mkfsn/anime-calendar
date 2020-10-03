<script>
    import calendarize from "calendarize";
    import CalendarDate from "./Calendar/CalendarDate.svelte";

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
        flex: 1 1;
        font-variant-numeric: tabular-nums;
    }
</style>

<div class="calendar">
    <div class="container">
        {#each { length:6 } as week, week_idx (week_idx)}
            {#if curr[week_idx]}
                <div class="week">
                    {#each { length:7 } as day, day_idx (day_idx)}
                        {#if curr[week_idx][day_idx] !== 0}
                            <CalendarDate day={curr[week_idx][day_idx]} isCurr={true} month={month} />
                        {:else if (week_idx < 1)}
                            <CalendarDate day={prev[prev.length - 1][day_idx]} isCurr={false} />
                        {:else}
                            <CalendarDate day={next[0][day_idx]} isCurr={false} />
                        {/if}
                    {/each}
                </div>
            {/if}
        {/each}
    </div>
</div>