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
      <div className="result-card result-empty">
        <div className="empty-icon">🔍</div>
        <p className="empty-text">Hasil parsing akan muncul di sini</p>
        <p className="empty-sub">Masukkan string dan klik Parse</p>
      </div>
    );
  }

  const fields = [
    {
      label: "NAME",
      value: result.raw.name,
      fmtValue: result.name,
      icon: "👤",
      color: "blue",
      width: 30,
    },
    {
      label: "AGE",
      value: result.raw.age,
      fmtValue: result.age,
      icon: "🎂",
      color: "purple",
      width: 3,
    },
    {
      label: "CITY",
      value: result.raw.city,
      fmtValue: result.city,
      icon: "🏙️",
      color: "teal",
      width: 20,
    },
  ];

  return (
    <div className="result-card result-success">
      <div className="result-header">
        <span className="result-badge">✓ Parsed</span>
        <span className="result-id">ID #{result.id}</span>
      </div>

      <div className="result-fields">
        {fields.map((f) => (
          <div key={f.label} className={`field-row field-${f.color}`}>
            <div className="field-icon">{f.icon}</div>
            <div className="field-content">
              <div className="field-label">{f.label}</div>
              <div className="field-value">{f.value || "—"}</div>
              <div className="field-meta">
                <span className="field-fmt-label">fixed({f.width})</span>
                <code className="field-fmt-value">&ldquo;{f.fmtValue}&rdquo;</code>
              </div>
            </div>
          </div>
        ))}
      </div>

      <div className="result-raw">
        <p className="raw-label">📋 Fixed-Width Output</p>
        <pre className="raw-block">
          <span className="raw-line">
            <span className="raw-key">NAME</span> : &ldquo;{result.name}&rdquo;
          </span>
          <span className="raw-line">
            <span className="raw-key">AGE </span> : &ldquo;{result.age}&rdquo;
          </span>
          <span className="raw-line">
            <span className="raw-key">CITY</span> : &ldquo;{result.city}&rdquo;
          </span>
        </pre>
      </div>
    </div>
  );
}
