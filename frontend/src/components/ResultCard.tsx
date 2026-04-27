"use client";

interface Props {
  result: {
    id: number;
    name: string;
    age: string;
    city: string;
    raw: { name: string; age: string; city: string };
    created_at: string;
  } | null;
}

export default function ResultCard({ result }: Props) {
  if (!result) {
    return (
      <div className="card card--muted">
        <p className="empty-text">Hasil parsing akan muncul di sini.</p>
      </div>
    );
  }

  return (
    <div className="card">
      <div className="result-header">
        <span className="result-title">Hasil parse</span>
        <span className="result-id">#{result.id}</span>
      </div>

      <table className="result-table">
        <thead>
          <tr>
            <th>Field</th>
            <th>Value</th>
            <th>Fixed-width ({`char`})</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>Name</td>
            <td>{result.raw.name || "—"}</td>
            <td><code>&ldquo;{result.name}&rdquo;</code></td>
          </tr>
          <tr>
            <td>Age</td>
            <td>{result.raw.age || "—"}</td>
            <td><code>&ldquo;{result.age}&rdquo;</code></td>
          </tr>
          <tr>
            <td>City</td>
            <td>{result.raw.city || "—"}</td>
            <td><code>&ldquo;{result.city}&rdquo;</code></td>
          </tr>
        </tbody>
      </table>
    </div>
  );
}