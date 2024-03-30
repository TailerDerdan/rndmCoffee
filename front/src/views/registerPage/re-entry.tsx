import React, { useState } from "react";
import stylesRegister from "./register.module.css";
import stylesReEntry from "./re-entry.module.css";
import { InputField } from "./register";
import { Navigate, useNavigate } from "react-router-dom";
import { ButtonTextLink } from "../../components/button/button";

export const ReEntryScreen = () => {
	const [InputNameEmail, setInputEmail] = useState("");
	const [InputNamePassword, setInputPassword] = useState("");
	const [InputErrorEmail, setInputErrorEmail] = useState(false);
	const [InputErrorPassword, setInputErrorPassword] = useState(false);
	const [ClassesError, setClassesError] = useState(
		stylesRegister.hidden__block +
			" " +
			stylesRegister.register__window__error,
	);
	const navigate = useNavigate();

	return (
		<div>
			<div className={stylesRegister.firstHalf}></div>
			<div className={stylesRegister.register}>
				<div className={ClassesError}>
					<div className={stylesRegister.wrapper__header__error}>
						<h1
							className={
								stylesRegister.register__window__error__header
							}
						>
							Такого пользователя не существует
						</h1>
					</div>
				</div>
				<div className={stylesRegister.register__window}>
					<div className={stylesRegister.window__header}>
						<h1 className={stylesRegister.register__header}>
							Вход
						</h1>
					</div>
					<form
						onSubmit={async (event) => {
							event.preventDefault();

							const data = {
								username: InputNameEmail,
								password: InputNamePassword,
							};

							if (data.username.length == 0) {
								setInputErrorEmail(true);
								return;
							}

							if (data.password.length == 0) {
								setInputErrorPassword(true);
								return;
							}

							const response = await fetch(
								"http://localhost:8000/auth/sign-in",
								{
									method: "POST",
									body: JSON.stringify(data),
									headers: {
										"Content-Type": "application/json",
									},
									credentials: "include",
								},
							);
							if (!response.ok) {
								if (response.status == 500) {
									setClassesError(
										stylesRegister.register__window__error,
									);
								}
								return;
							}
							navigate("/main", { replace: true });
						}}
						action="http:://localhost:8000/auth/sign-in"
						method="post"
					>
						<div>
							<InputField
								InputData={InputNameEmail}
								setInput={setInputEmail}
								title={"Email"}
								type={"email"}
								autoComplete={true}
								id={"username"}
								error={InputErrorEmail}
								setErrorInput={setInputErrorEmail}
								setErrorData={setClassesError}
							/>
							<InputField
								InputData={InputNamePassword}
								setInput={setInputPassword}
								title={"Пароль"}
								type={"password"}
								autoComplete={true}
								id={"password"}
								error={InputErrorPassword}
								setErrorInput={setInputErrorPassword}
								setErrorData={setClassesError}
							/>
						</div>
						<div
							className={stylesReEntry.wrapper__user__assistance}
						>
							<div
								className={
									stylesReEntry.wrapper__remember__entry__input
								}
							>
								<input
									type={"checkbox"}
									className={
										stylesReEntry.remember__entry__input
									}
								/>
								<label
									className={
										stylesReEntry.label__remember__entry__input
									}
								>
									Запомнить вход
								</label>
							</div>
							<div
								className={
									stylesReEntry.wrapper__forgot__password
								}
							>
								<a
									className={stylesReEntry.forgot__password}
									href={"#"}
								>
									Забыли пароль?
								</a>
							</div>
						</div>
						<div className={stylesRegister.wrapper__input__submit}>
							<button
								type={"submit"}
								className={stylesRegister.input__submit}
							>
								Вход
							</button>
						</div>
						<div className={stylesReEntry.wrapper__create__account}>
							<div>
								<h2
									className={
										stylesRegister.inputField__header
									}
								>
									Еще не с нами?
								</h2>
							</div>
							<div
								className={
									stylesRegister.wrapper__buttonTextLink
								}
							>
								<ButtonTextLink
									id={"reg"}
									title={"Регистрация"}
									link={"/reg"}
								/>
							</div>
						</div>
					</form>
				</div>
			</div>
		</div>
	);
};
export const foo = () => {
	return <></>;
};
