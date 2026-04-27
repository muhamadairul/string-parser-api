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
    <div className="history-card">
      <div className="history-header">
        <div className="history-title-row">
          <h3 className="history-title">🗂 Parse History</h3>
          <span className="history-count">{history.length} records</span>
        </div>
        <button className="refresh-btn" onClick={fetchHistory} disabled={loading}>
          {loading ? "⟳" : "↻"} Refresh
        </button>
      </div>

      {history.length === 0 && !loading ? (
        <div className="history-empty">
          <p>Belum ada data. Coba parse string pertama kamu!</p>
        </div>
      ) : (
        <div className="table-wrapper">
          <table className="history-table">
            <thead>
              <tr>
                <th>#</th>
                <th>NAME (30)</th>
                <th>AGE (3)</th>
                <th>CITY (20)</th>
                <th>WAKTU</th>
              </tr>
            </thead>
            <tbody>
              {history.map((item, idx) => (
                <tr key={item.id} className="table-row">
                  <td className="td-id">{idx + 1}</td>
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
