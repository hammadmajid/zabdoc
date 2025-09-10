import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

export const POST: RequestHandler = async ({ request, url }) => {
    try {
        const body = await request.formData();

        let base = "http://localhost:8080";

        // Shitty hack: change hostname to zabdoc.me in prod
        if (url.hostname === "zabdoc.me") {
            base = "https://api.zabdoc.me"
        }

        const response = await fetch(`${base}/assignment`, {
            method: "POST",
            body: body,
        });

        if (!response.ok) {
            throw "Internal server error";
        }

        const pdfBuffer = await response.arrayBuffer();

        return new Response(pdfBuffer, {
            status: 200,
            headers: {
                "Content-Type": "application/pdf",
                "Content-Disposition": "attachment; filename=assignment.pdf"
            }
        });
    } catch (error) {
        console.log(error)
        return json(
            { error: 'Internal server error' },
            { status: 500 }
        );
    }
};