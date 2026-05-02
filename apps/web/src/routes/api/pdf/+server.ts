// src/routes/api/pdf/+server.ts
import { json, type RequestHandler } from '@sveltejs/kit';
import { buildHTML } from '$lib/html';
import { documentSchema } from '@repo/types';
import { CLOUDFLARE_ACCOUNT_ID, CLOUDFLARE_API_TOKEN } from "$env/static/private";

export const POST: RequestHandler = async ({ request }) => {
  try {
    // Parse JSON from req
    const jsonData = await request.json();

    // Validate with schema
    const validated = documentSchema.parse(jsonData);

    // Ensure required fields have defaults
    const validatedWithDefaults = {
      ...validated,
      Students: validated.Students.map(student => ({
        Name: student.Name || '',
        RegNo: student.RegNo || '',
      })),
    };

    // Build HTML
    const html = buildHTML(validatedWithDefaults);

    // Send to Cloudflare
    const response = await fetch(
      `https://api.cloudflare.com/client/v4/accounts/${CLOUDFLARE_ACCOUNT_ID}/browser-rendering/pdf`,
      {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${CLOUDFLARE_API_TOKEN}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ html })
      }
    );

    if (!response.ok) {
      const errorData = await response.json();
      return json({ error: 'Cloudflare API failed', details: errorData }, { status: response.status });
    }

    const pdfBuffer = Buffer.from(await response.arrayBuffer());

    return new Response(pdfBuffer, {
      headers: {
        'Content-Type': 'application/pdf',
        'Content-Length': pdfBuffer.length.toString(),
        'Cache-Control': 'public, max-age=3600',
      },
    });
  } catch (error) {
    console.error('PDF generation error:', error);

    // Return validation error if zod failed
    if (error instanceof Error && error.name === 'ZodError') {
      return json(
        { error: 'Invalid data', details: error.message },
        { status: 400 }
      );
    }

    return json({ error: 'Internal server error' }, { status: 500 });
  }
};