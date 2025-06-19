export default function TopTips() {
  return (
    <div className="bg-white dark:bg-gray-800 p-4 rounded-lg shadow-md">
      <h2 className="text-lg font-semibold mb-2">Top Tips</h2>
      <ul className="list-disc pl-5 space-y-2">
        <li>Use the search bar to quickly find entities.</li>
        <li>Click on an entity to view its details.</li>
        <li>Use filters to narrow down your search results.</li>
        <li>Bookmark frequently accessed entities for quick access.</li>
      </ul>
    </div>
  );
}
