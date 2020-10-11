<script>
    export let entries;
</script>

<table>
    {#each Array.from(entries) as [k, v] }
        <tr>
            <th>{k}</th>
            <td>{v}</td>
        </tr>
    {/each}
</table>