<script lang="ts">
  import { changelog, formatDate } from "$lib/changelog";
</script>

<div class="space-y-8">
    <div class="space-y-2">
        <h1 class="text-4xl font-black uppercase tracking-tight">Changelog</h1>
        <p class="text-muted-foreground">
            History of zabdoc versions and improvements
        </p>
    </div>

    <div class="space-y-6">
        {#each changelog as entry, index (entry.version)}
            <div class="neo-border-sm neo-shadow-sm p-6 space-y-3 relative">
                {#if index === 0}
                    <div class="absolute -top-3 -right-3 bg-primary px-3 py-1 neo-border-sm text-xs font-black uppercase tracking-tight">
                        Latest
                    </div>
                {/if}
                <div class="flex items-baseline gap-3">
                    <h2 class="text-2xl font-black uppercase tracking-tight">
                        v{entry.version}
                    </h2>
                    <span class="text-sm font-medium text-muted-foreground">
                        {formatDate(entry.date)}
                    </span>
                </div>

                <div class="space-y-2">
                    <h3 class="text-sm font-bold uppercase tracking-wide text-muted-foreground">
                        Changes
                    </h3>
                    <ul class="space-y-1">
                        {#each entry.changes as change}
                            <li class="text-sm leading-relaxed">
                                <span class="font-bold">•</span>
                                {change}
                            </li>
                        {/each}
                    </ul>
                </div>
            </div>
        {/each}
    </div>
</div>
