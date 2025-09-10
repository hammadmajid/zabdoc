// src/routes/api/users/+server.ts
import { json } from '@sveltejs/kit';
import type { RequestHandler } from '$types';

export const POST: RequestHandler = async ({ request }) => {
    try {
        const body = await request.formData();

        const response = await fetch("http://localhost:8080/assignment", {
            method: "POST",
            headers: {
                "Content-Type": "multipart/form-data"
            },
            body: body, // stringify the parsed JSON object
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