<script lang="ts">
    interface ChangelogEntry {
        version: string;
        date: string;
        changes: string[];
    }

    function formatDate(dateString: string): string {
        const date = new Date(dateString + "T00:00:00");
        return new Intl.DateTimeFormat("en-US", {
            year: "numeric",
            month: "short",
            day: "numeric"
        }).format(date);
    }

    const changelog: ChangelogEntry[] = [
        {
            version: "2.2.1",
            date: "2026-02-20",
            changes: [
                "Only show relevant courses based on document type",
                "Added logging info in privacy policy",
                "Fixed storage card width on mobile devices",
                "Removed description field from changelog"
            ]
        },
        {
            version: "2.2.0",
            date: "2026-02-15",
            changes: [
                "Added changelog page for version history",
                "Render unique icon for each course",
                "Fixed settings storage card functionality",
                "Improved overall UX with bug fixes",
                "Restricted assignment number selector range to 1-4"
            ]
        },
        {
            version: "2.1.1",
            date: "2026-02-15",
            changes: [
                "Fixed about page styling",
                "Updated version display in navbar",
                "Improved radio group component for better UX",
                "Fixed wizard state management with abort controller",
                "Updated assignment number selector",
                "Added data for spring 2026"
            ]
        },
        {
            version: "2.1.0",
            date: "2026-02-14",
            changes: [
                "Added radio group component for better selection UX",
                "Implemented fixed number of assignment selector",
                "Improved wizard component interactions",
                "Enhanced data handling for assignments",
                "Better state management across wizard steps"
            ]
        },
        {
            version: "2.0.0",
            date: "2026-01-17",
            changes: [
                "Implemented neobrutalist design language UI",
                "Introduced wizard component for intuitive document generation",
                "Redesigned favicon for consistency",
                "Restructured UI elements for better user experience",
                "Removed outdated pages and components"
            ]
        },
        {
            version: "1.0.0",
            date: "2025-09-09",
            changes: [
                "Migrated to SvelteKit for better performance",
                "Implemented touch drag and drop functionality for mobile",
                "Modernized tech stack and tooling",
                "Improved code organization and maintainability",
                "Better component architecture"
            ]
        },
        {
            version: "0.1.0",
            date: "2025-09-02",
            changes: [
                "Initial project setup",
                "Basic document generation functionality",
                "GNU General Public License v3 adoption"
            ]
        }
    ];
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
