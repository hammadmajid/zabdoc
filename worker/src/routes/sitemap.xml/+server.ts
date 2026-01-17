export async function GET() {
    return new Response(
        `
		<?xml version="1.0" encoding="UTF-8"?>
        <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
            <url>
                <loc>https://zabdoc.xyz/</loc>
                <lastmod>2025-09-20</lastmod>
                <changefreq>weekly</changefreq>
                <priority>1.0</priority>
            </url>
            <url>
                <loc>https://zabdoc.xyz/about</loc>
                <lastmod>2025-09-20</lastmod>
                <changefreq>weekly</changefreq>
                <priority>0.8</priority>
            </url>
            <url>
                <loc>https://zabdoc.xyz/settings</loc>
                <lastmod>2025-09-20</lastmod>
                <changefreq>weekly</changefreq>
                <priority>0.6</priority>
            </url>
        </urlset>`.trim(),
        {
            headers: {
                'Content-Type': 'application/xml'
            }
        }
    );
}