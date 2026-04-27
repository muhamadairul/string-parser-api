"use client";

import { useState, useEffect, useCallback } from "react";

interface HistoryItem {
  id: number;
  name: string;
  age: string;
  city: string;
  created_at: string;
}

interface Props {
  refreshTrigger: number;
}

export default function HistoryTable({ refreshTrigger }: Props) {
  const [history, setHistory] = useState<HistoryItem[]>([]);
  const [loading, setLoading] = useState(false);

  const fetchHistory = useCallback(async () => {
    setLoading(true);
    try {
      const res = await fetch("http://localhost:8080/api/history");
      const json = await res.json();
      if (json.status === "success" && Array.isArray(json.data)) {
        setHistory(json.data);
      }
    } catch {
      // silently fail
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    fetchHistory();
  }, [fetchHistory, refreshTrigger]);

  const formatDate = (dateStr: string) => {
    try {
      return new Date(dateStr).toLocaleString("id-ID", {
        day: "2-digit",
        month: "short",
        year: "numeric",
        hour: "2-digit",
        minute: "2-digit",
      });
    } catch {
      return dateStr;
    }
  };

  return (
    <div className="card">
      <div className="history-header">
        <span className="history-title">History</span>
        <div className="history-meta">
          <span className="history-count">{history.length} record</span>
          <button className="btn-ghost" onClick={fetchHistory} disabled={loading}>
            {loading ? "..." : "Refresh"}
          </button>
        </div>
      </div>

      {history.length === 0 && !loading ? (
        <p className="empty-text">Belum ada data.</p>
      ) : (
        <div className="table-wrapper">
          <table className="data-table">
            <thead>
              <tr>
                <th>#</th>
                <th>Name (30)</th>
                <th>Age (3)</th>
                <th>City (20)</th>
                <th>Waktu</th>
              </tr>
            </thead>
            <tbody>
              {history.map((item, idx) => (
                <tr key={item.id}>
                  <td className="td-num">{idx + 1}</td>
                  <td className="td-mono">{item.name.trim()}</td>
                  <td className="td-mono td-center">{item.age.trim()}</td>
                  <td className="td-mono">{item.city.trim()}</td>
                  <td className="td-time">{formatDate(item.created_at)}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}