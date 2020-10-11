<script>
    export let keys;
    export let values;
    export let showKey;
    export let valueToRows;
    export let highlightValue = () => '';
    let selected;
</script>

<select bind:value={selected}>
    {#each keys as key, index}
        <option value={index}>
            {#if showKey}
                {showKey(key)}
            {:else}
                {key}
            {/if}
        </option>
    {/each}
</select>

{#if selected !== undefined}
    <table>
        <tbody>
        {#each values[selected] as value}
            <tr class:active={highlightValue(value)}>
                {#if valueToRows}
                    {#each valueToRows(value) as row}
                        <td>{row}</td>
                    {/each}
                {/if}
            </tr>
        {/each}
        </tbody>
    </table>
{/if}

<style>
    tr.active {
        background-color: yellow;
    }
</style>