interface ChangelogEntry {
    version: string;
    date: string;
    changes: string[];
}

export function formatDate(dateString: string): string {
    const date = new Date(dateString + "T00:00:00");
    return new Intl.DateTimeFormat("en-US", {
        year: "numeric",
        month: "short",
        day: "numeric"
    }).format(date);
}

export const changelog: ChangelogEntry[] = [
    {
        version: "3.0.2",
        date: "2026-03-14",
        changes: [
            "Fixed spelling from 'Scrap' to 'Scrape' in data-related sections",
            "Updated about page content",
            "Added all courses to the dataset",
            "Added script to clean data",
            "Removed course icons"
        ]
    },
    {
        version: "3.0.1",
        date: "2026-03-11",
        changes: [
            "Redesigned tabs UI with neobrualist style and purple accent color",
            "Redesigned table component with bold borders and uppercase headers",
            "Implemented proper Table components in scrap page",
            "Fixed table row borders for consistent styling",
            "Fixed back to home button visibility on document generation page",
            "Reduced heading sizes across pages",
            "Removed reset button from scrap page"
        ]
    },
    {
        version: "3.0.0",
        date: "2026-03-11",
        changes: [
            "Added attendance scraping feature for ZabDesk integration",
            "Introduced new /scrap route for fetching attendance data",
            "Moved document generation wizard to dedicated /document route",
            "Redesigned home page with choice between Scrap Data and Generate Document",
            "Implemented attendance display with tabbed interface (Attendance/Marks)",
            "Added color-coded attendance status badges (Present/Absent/Late)"
        ]
    },
    {
        version: "2.2.2",
        date: "2026-02-21",
        changes: [
            "Fixed $effect error in form store for course filtering",
            "Improved course validation when document type changes",
            "Enhanced reactive course list updates"
        ]
    },
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