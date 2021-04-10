# controllers/auth_controller.ex
defmodule NorzeWeb.AuthController do
  @moduledoc false

  use NorzeWeb, :controller
  plug Ueberauth

  alias Norze.Accounts

  def delete(conn, _params) do
    conn
    |> configure_session(drop: true)
    |> redirect(to: "/")
  end

  def callback(%{assigns: %{ueberauth_failure: _fails}} = conn, _params) do
    conn
    |> put_flash(:error, "Failed to authenticate.")
    |> redirect(to: "/")
  end

  def callback(%{assigns: %{ueberauth_auth: auth}} = conn, _params) do
    case Accounts.from_auth(auth) do
      {:ok, user} ->
        conn
        |> put_session(:current_user_id, user.id)
        |> configure_session(renew: true)
        |> redirect(to: "/")

      {:error, _} ->
        conn
        |> put_flash(:error, "Failed authenticated.")
        |> redirect(to: "/")
    end
  end
end
