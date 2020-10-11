<script>
    import {Tabs, Tab, TabList, TabPanel} from 'svelte-tabs';
    import Table from "../Table.svelte";
    import DropdownTabs from "../DropdownTabs.svelte";

    export let anime;
    export let episode;

    function toHighlight(_episode) {
        return _episode === episode;
    }

    function backgroundImage(anime) {
        if (anime.OpenGraph && anime.OpenGraph.Images && anime.OpenGraph.Images.length > 0) {
            return `url(${anime.OpenGraph.Images[0].URL})`
        }
        return "none"
    }
</script>

<div class="container" style="--background-image: {backgroundImage(anime)}">
    <h1>
        {anime.Title}
        {#if episode}
            <small>
                ({episode.Number}) {episode.Subtitle}
            </small>
        {/if}
    </h1>

    <Tabs>
        <TabList>
            {#if anime.Cast}
                <Tab>Cast</Tab>
            {/if}
            {#if anime.Staff}
                <Tab>Staff</Tab>
            {/if}
            {#if anime.Channels && anime.Channels.length != 0}
                <Tab>Channels</Tab>
            {/if}
        </TabList>

        {#if anime.Cast}
            <TabPanel>
                <section>
                    <Table entries={Object.entries(anime.Cast)} />
                </section>
            </TabPanel>
        {/if}

        {#if anime.Staff}
            <TabPanel>
                <section>
                    <Table entries={Object.entries(anime.Staff)} />
                </section>
            </TabPanel>
        {/if}

        {#if anime.Channels && anime.Channels.length != 0}
            <TabPanel>
                <section>
                    <div class="content">
                        <DropdownTabs
                            highlightValue={toHighlight}
                            keys={anime.Channels}
                            values={anime.Timetables}
                            valueToRows={(v) => [v.Number, v.StartAt, v.Subtitle]}
                        />
                    </div>
                </section>
            </TabPanel>
        {/if}

    </Tabs>
</div>

<style>
    h1 {
        z-index: 2;
        position: relative;
    }
    h1 > small {
        display: block;
        color: #7f7f7f;
        font-size: .7em;
    }
    .container {
        padding: 1em;
        position: relative;
    }
    .container::after {
        content: "";
        background-image: var(--background-image);
        background-size: cover;
        background-repeat: no-repeat;
        background-position: center;
        position: absolute;
        opacity: 0.15;
        top: 0;
        left: 0;
        bottom: 0;
        right: 0;
        z-index: 1;
    }
    .container :global(.svelte-tabs) {
        position: relative;
        z-index: 2;
    }
    section {
        display: flex;
        justify-content: center;
    }
    .container :global(.svelte-tabs .svelte-tabs__tab-list) {
        border-bottom: none;
    }
</style>

