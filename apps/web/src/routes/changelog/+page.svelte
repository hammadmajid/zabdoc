<script lang="ts">
	import { changelog, formatDate } from "$lib/changelog";
</script>

<div class="space-y-8">
	<div class="space-y-2">
		<h1 class="text-4xl font-black tracking-tight uppercase">Changelog</h1>
		<p class="text-muted-foreground">History of zabdoc versions and improvements</p>
	</div>

	<div class="space-y-6">
		{#each changelog as entry, index (entry.version)}
			<div class="neo-border-sm neo-shadow-sm relative space-y-3 p-6">
				{#if index === 0}
					<div
						class="neo-border-sm absolute -top-3 -right-3 bg-primary px-3 py-1 text-xs font-black tracking-tight uppercase"
					>
						Latest
					</div>
				{/if}
				<div class="flex items-baseline gap-3">
					<h2 class="text-2xl font-black tracking-tight uppercase">
						v{entry.version}
					</h2>
					<span class="text-sm font-medium text-muted-foreground">
						{formatDate(entry.date)}
					</span>
				</div>

				<div class="space-y-2">
					<h3 class="text-sm font-bold tracking-wide text-muted-foreground uppercase">
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
