<script>
    import {calendarStore} from "../../stores";

    function formatter(d) {
        const year = d.getFullYear();
        const month = ('0' + (d.getMonth()+1)).slice(-2);
        return `${year}年${month}月`
    }

    function toPrev() {
        calendarStore.update(d => new Date(d.setMonth(d.getMonth()-1)))
    }

    function toNext() {
        calendarStore.update(d => new Date(d.setMonth(d.getMonth()+1)))
    }
</script>

<div class="controller">
    <p class="column arrow" on:click={toPrev}>
        <svg class="svg-icon" viewBox="0 0 20 20">
            <path d="M11.739,13.962c-0.087,0.086-0.199,0.131-0.312,0.131c-0.112,0-0.226-0.045-0.312-0.131l-3.738-3.736c-0.173-0.173-0.173-0.454,0-0.626l3.559-3.562c0.173-0.175,0.454-0.173,0.626,0c0.173,0.172,0.173,0.451,0,0.624l-3.248,3.25l3.425,3.426C11.911,13.511,11.911,13.789,11.739,13.962 M18.406,10c0,4.644-3.763,8.406-8.406,8.406S1.594,14.644,1.594,10S5.356,1.594,10,1.594S18.406,5.356,18.406,10 M17.521,10c0-4.148-3.373-7.521-7.521-7.521c-4.148,0-7.521,3.374-7.521,7.521c0,4.148,3.374,7.521,7.521,7.521C14.147,17.521,17.521,14.148,17.521,10"></path>
        </svg>
    </p>
    <p class="column date">{formatter($calendarStore)}</p>
    <p class="column arrow" on:click={toNext}>
        <svg class="svg-icon" viewBox="0 0 20 20">
            <path d="M12.522,10.4l-3.559,3.562c-0.172,0.173-0.451,0.176-0.625,0c-0.173-0.173-0.173-0.451,0-0.624l3.248-3.25L8.161,6.662c-0.173-0.173-0.173-0.452,0-0.624c0.172-0.175,0.451-0.175,0.624,0l3.738,3.736C12.695,9.947,12.695,10.228,12.522,10.4 M18.406,10c0,4.644-3.764,8.406-8.406,8.406c-4.644,0-8.406-3.763-8.406-8.406S5.356,1.594,10,1.594C14.643,1.594,18.406,5.356,18.406,10M17.521,10c0-4.148-3.374-7.521-7.521-7.521c-4.148,0-7.521,3.374-7.521,7.521c0,4.147,3.374,7.521,7.521,7.521C14.147,17.521,17.521,14.147,17.521,10"></path>
        </svg>
    </p>
</div>

<style>
    .column {
        display: inline-block;
    }
    .date {
        font-variant-numeric: tabular-nums;
        margin: 0;
    }
    .arrow {
        margin: 0;
        vertical-align: bottom;
    }
    .arrow svg {
        border-radius: 50%;
        width: 1.5em;
        height: 1.5em;
        display: block;
        text-align: center;
    }
    .arrow svg:hover {
        background-color: #d0d0d0;
        cursor: pointer;
    }
</style>
