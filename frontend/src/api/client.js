export async function apiFetch(endpoint, options = {}) {
  // ローカルストレージに保存しているトークンを取得
  const token = localStorage.getItem("token");

  try {
    //apiをたたく
    const response = await fetch(endpoint, {
      ...options,
      headers: {
        "Content-Type": "application/json",
        ...(token && { Authorization: "Bearer ${token}" }), //tokenがある場合
        ...options.headers,
      },
    });

    if (!response.ok) {
      throw new Error("取得に失敗しました");
    }
    return await response.json();
  } catch (error) {
    console.error(error);
    throw error;
  }
}
