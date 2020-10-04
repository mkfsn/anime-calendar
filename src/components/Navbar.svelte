<script>
    import CalendarController from "./Calendar/CalendarController.svelte";
    export let title = "Anime Bangumi";
    export let active;
    function toggleActive() {
        active = !active;
    }
</script>

<style>
    .navbar {
        height: 3em;
        border-bottom: 1px solid #cccccc;
        padding: .5em;
        display: flex;
    }
    .brand {
        display: flex;
        align-items: center;
    }
    .brand > i {
        border-radius: 50%;
        width: 3em;
        height: 3em;
        text-align: center;
        line-height: 3em;
        display: flex;
        align-items: inherit;
        justify-content: center;
        margin: 0 .5em;
    }
    .brand > i:hover {
        background-color: #efefef;
        cursor: pointer;
    }
    .brand > i > svg {
        width: 60%;
        height: 60%;
    }
    .brand > i > svg > path {
        fill: gray;
    }
    .brand > i.active {
        background-color: #e0e0e0;
    }
    .brand h1 {
        margin: 0;
        font-size: 1.3em;
    }
    .left {
        flex: 1 1;
    }
    .right {
        flex: 1 1;
        text-align: right;
        align-self: center;
    }
</style>


<div class="navbar">
    <div class="left">
        <div class="brand">
            <i class:active={active} on:click={toggleActive}>
                <svg class="svg-icon" viewBox="0 0 20 20">
                    <path fill="none" d="M3.314,4.8h13.372c0.41,0,0.743-0.333,0.743-0.743c0-0.41-0.333-0.743-0.743-0.743H3.314
			    					c-0.41,0-0.743,0.333-0.743,0.743C2.571,4.467,2.904,4.8,3.314,4.8z M16.686,15.2H3.314c-0.41,0-0.743,0.333-0.743,0.743
			    					s0.333,0.743,0.743,0.743h13.372c0.41,0,0.743-0.333,0.743-0.743S17.096,15.2,16.686,15.2z M16.686,9.257H3.314
			    					c-0.41,0-0.743,0.333-0.743,0.743s0.333,0.743,0.743,0.743h13.372c0.41,0,0.743-0.333,0.743-0.743S17.096,9.257,16.686,9.257z"></path>
                </svg>
            </i>
            <h1>{ title }</h1>
        </div>
    </div>
    <div class="right">
        <CalendarController />
    </div>
</div>