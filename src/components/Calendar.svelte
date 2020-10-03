<script>
    import calendarize from "calendarize";
    import CalendarDate from "./Calendar/CalendarDate.svelte";
    import {dateStore} from "../store";

    let thisMonth = $dateStore,
        prevMonth = new Date(thisMonth),
        nextMonth = new Date(thisMonth);
    prevMonth.setMonth(prevMonth.getMonth()-1);
    nextMonth.setMonth(nextMonth.getMonth()+1);

    let prev = calendarize(prevMonth),
        curr = calendarize(thisMonth),
        next = calendarize(nextMonth);

    dateStore.subscribe(d => {
        thisMonth = d;
        prevMonth = new Date(d);
        nextMonth = new Date(d);
        prevMonth.setMonth(prevMonth.getMonth()-1);
        nextMonth.setMonth(nextMonth.getMonth()+1);
        console.log({thisMonth, prevMonth, nextMonth});
        prev = calendarize(prevMonth);
        curr = calendarize(thisMonth);
        next = calendarize(nextMonth);
        console.log({prev, curr, next});
    })
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
        {#each { length:6 } as _, week_idx (week_idx)}
            {#if curr[week_idx]}
                <div class="week">
                    {#each { length:7 } as _, day_idx (day_idx)}
                        {#if curr[week_idx][day_idx] !== 0}
                            <CalendarDate day={curr[week_idx][day_idx]} isCurr={true} date={thisMonth} />
                        {:else if (week_idx < 1)}
                            <CalendarDate day={prev[prev.length - 1][day_idx]} isCurr={false} date={prevMonth} />
                        {:else}
                            <CalendarDate day={next[0][day_idx]} isCurr={false} date={nextMonth} />
                        {/if}
                    {/each}
                </div>
            {/if}
        {/each}
    </div>
</div>