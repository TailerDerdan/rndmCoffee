import React, { useEffect, useRef, useState } from "react";
import stylesProfile from "./mainProfilePage.module.css";
import stylesHomePage from "../homePage/homePage.module.css";
import stylesContacts from "../../profilePage/contacts/contactsPage.module.css";
import { ManInChairIcon, ProfileIcon } from "../../../components/icons/icons";
import {
	LocationInputField,
	TextField,
	TypeInputOnProfile,
} from "../../../components/input/input";
import { useLocalStorage } from "../../../hooks/useLocalStorage";
import { Button, ButtonType } from "../../../components/button/button";
import { ChoiceActivity } from "../../../components/choiceActivity/choiceActivity";
import { TypeOfActivities } from "../../../data/types";
import { AllActivities } from "../../../data/data";

export const MainProfilePage = () => {
	const [nameUser, setNameUser] = useState("");
	const [nameUserError, setNameUserError] = useState(false);

	const [surnameUser, setSurNameUser] = useState("");
	const [surnameUserError, setSurnameUserError] = useState(false);

	const [emailUser, setEmailUser] = useState("");
	const [emailUserError, setEmailUserError] = useState(false);

	const [cityUser, setCityUser] = useState("");

	const [birthdayUser, setBirthdayUser] = useState("");
	const [description, setActivities] = useState(Array<TypeOfActivities>);

	const [classesForContactInfo, setClassesForContactInfo] = useState(
		stylesProfile.coctactInfo + " " + stylesProfile.noClick,
	);
	const [classesForButtonEdit, setClassesForButtonEdit] = useState(
		stylesProfile.wrapper__buttonEdit,
	);
	const [classesForButtonClose, setClassesForButtonClose] = useState(
		stylesProfile.button__close + " " + stylesProfile.invisibleBlock,
	);
	const [classesForButtonSave, setClassesForButtonSave] = useState(
		stylesProfile.button__save + " " + stylesProfile.invisibleBlock,
	);
	const [edit, setEdit] = useState(false);

	const refContactInfo = useRef<HTMLDivElement>(null);

	const [token] = useLocalStorage({
		initialValue: "",
		key: "token",
	});

	const [profile_id] = useLocalStorage({
		initialValue: -1,
		key: "profile_id",
	});

	const [username, setUsername] = useLocalStorage({
		initialValue: "",
		key: "username",
	});

	useEffect(() => {
		async function getProfile() {
			const res = await fetch(
				`http://localhost:8000/api/profile/get_profile/${profile_id}`,
				{
					method: "GET",
					headers: {
						"Content-Type": "application/json",
						Authorization: `Bearer ${token}`,
					},
					credentials: "include",
				},
			);
			if (res.ok) {
				const data = await res.json();
				setNameUser(data.data.name);
				setEmailUser(data.data.email);
				setSurNameUser(data.data.surname);
				setBirthdayUser(data.data.birthday.substr(0, 10));
				setCityUser(data.data.city);
			}
		}
		async function getHobbies() {
			const res = await fetch(
				`http://localhost:8000/api/profile/${profile_id}/hobby/get_hobby`,
				{
					method: "GET",
					headers: {
						"Content-Type": "application/json",
						Authorization: `Bearer ${token}`,
					},
					credentials: "include",
				},
			);
			if (res.ok) {
				const data = await res.json();
				const hobbies: Array<TypeOfActivities> = [];
				if (data.data) {
					data.data.forEach(
						(hobby: { id: number; description: string }) => {
							hobbies.push({ description: hobby.description });
						},
					);
				}
				setActivities(hobbies);
			}
		}
		getProfile();
		getHobbies();
	}, []);

	return (
		<div className={stylesProfile.mainContent}>
			<div className={stylesHomePage.mainContent__iconMan}>
				<ManInChairIcon />
			</div>
			<div className={stylesProfile.mainContent__contactInfo}>
				<div>
					<Button
						id={"avatar"}
						type={ButtonType.Icon}
						icon={<ProfileIcon />}
					/>
				</div>
				<form
					onSubmit={async (event) => {
						event.preventDefault();

						const data = {
							city: cityUser,
							findstatus: true,
							name: nameUser,
							photo: "none",
							surname: surnameUser,
							birthday: birthdayUser,
							email: emailUser,
						};

						if (data.city.length == 0) {
							data.city = "none";
						}
						if (data.name.length == 0) {
							setNameUserError(true);
							return;
						}
						if (data.surname.length == 0) {
							setSurnameUserError(true);
							return;
						}
						if (data.email.length == 0) {
							setEmailUserError(true);
							return;
						}

						const response = await fetch(
							`http://localhost:8000/api/profile/edit_profile/${profile_id}`,
							{
								method: "PUT",
								body: JSON.stringify(data),
								headers: {
									"Content-Type": "application/json",
									Authorization: `Bearer ${token}`,
								},
								credentials: "include",
							},
						);
						if (response.ok) {
							setClassesForContactInfo(
								stylesProfile.coctactInfo +
									" " +
									stylesProfile.noClick,
							);
							setClassesForButtonEdit(
								stylesProfile.wrapper__buttonEdit,
							);
							setClassesForButtonClose(
								stylesProfile.button__close +
									" " +
									stylesProfile.invisibleBlock,
							);
							setClassesForButtonSave(
								stylesProfile.button__save +
									" " +
									stylesProfile.invisibleBlock,
							);
							setEdit(!edit);
							setUsername(nameUser);
						}
					}}
					action={`http://localhost:8000/api/profile/edit_profile/${profile_id}`}
					className={stylesProfile.contact__form}
				>
					<div className={classesForContactInfo} ref={refContactInfo}>
						<div className={stylesContacts.wrapper__inputs}>
							<div className={stylesContacts.wrapper__input__lfs}>
								{
									<TextField
										id={"name"}
										textLabel={"Имя"}
										typeInput={"text"}
										inputData={nameUser}
										setInput={setNameUser}
										location={
											LocationInputField.MainProfile
										}
										typeInputOnProfile={
											TypeInputOnProfile.Double
										}
										edit={edit}
										error={nameUserError}
										setErrorInput={setNameUserError}
									/>
								}
							</div>
							<div className={stylesContacts.wrapper__input__rfs}>
								{
									<TextField
										id={"surname"}
										textLabel={"Фамилия"}
										typeInput={"text"}
										inputData={surnameUser}
										setInput={setSurNameUser}
										location={
											LocationInputField.MainProfile
										}
										typeInputOnProfile={
											TypeInputOnProfile.Double
										}
										edit={edit}
										error={surnameUserError}
										setErrorInput={setSurnameUserError}
									/>
								}
							</div>
						</div>
						<div className={stylesContacts.wrapper__inputs}>
							{
								<TextField
									id={"email"}
									textLabel={"Email"}
									typeInput={"email"}
									inputData={emailUser}
									setInput={setEmailUser}
									location={LocationInputField.MainProfile}
									typeInputOnProfile={
										TypeInputOnProfile.Single
									}
									edit={edit}
									error={emailUserError}
									setErrorInput={setEmailUserError}
								/>
							}
						</div>
						<div className={stylesContacts.wrapper__inputs}>
							{
								<TextField
									id={"birthday"}
									textLabel={"Дата рождения"}
									typeInput={"date"}
									inputData={birthdayUser}
									setInput={setBirthdayUser}
									location={LocationInputField.MainProfile}
									typeInputOnProfile={
										TypeInputOnProfile.Single
									}
									edit={edit}
								/>
							}
						</div>
						<div className={stylesContacts.wrapper__inputs}>
							{
								<TextField
									id={"city"}
									textLabel={"Город"}
									typeInput={"text"}
									inputData={cityUser}
									setInput={setCityUser}
									location={LocationInputField.MainProfile}
									typeInputOnProfile={
										TypeInputOnProfile.Single
									}
									edit={edit}
								/>
							}
						</div>
						<div className={stylesProfile.wrapper__choiceActivity}>
							{
								<ChoiceActivity
									AllActivities={AllActivities}
									activities={description}
									setActivities={setActivities}
								/>
							}
						</div>
					</div>
					<div className={classesForButtonEdit}>
						<Button
							onClick={() => {
								setClassesForContactInfo(
									stylesProfile.coctactInfo,
								);
								setClassesForButtonEdit(
									stylesProfile.wrapper__buttonEdit +
										" " +
										stylesProfile.invisibleBlock,
								);
								setClassesForButtonClose(
									stylesProfile.button__close,
								);
								setClassesForButtonSave(
									stylesProfile.button__save,
								);
								setEdit(!edit);
							}}
							id={"editProfile"}
							type={ButtonType.Text}
							title={"Редактировать"}
						/>
					</div>
					<div className={stylesProfile.buttons__afterEdit}>
						<div className={classesForButtonClose}>
							<Button
								onClick={() => {
									setClassesForContactInfo(
										stylesProfile.coctactInfo +
											" " +
											stylesProfile.noClick,
									);
									setClassesForButtonEdit(
										stylesProfile.wrapper__buttonEdit,
									);
									setClassesForButtonClose(
										stylesProfile.button__close +
											" " +
											stylesProfile.invisibleBlock,
									);
									setClassesForButtonSave(
										stylesProfile.button__save +
											" " +
											stylesProfile.invisibleBlock,
									);
									setEdit(!edit);
								}}
								id={"closeEdit"}
								type={ButtonType.Text}
								title={"Закрыть"}
							/>
						</div>
						<div className={classesForButtonSave}>
							<Button
								typeButton={"submit"}
								id={"saveEdit"}
								type={ButtonType.Text}
								title={"Сохранить"}
							/>
						</div>
					</div>
				</form>
			</div>
		</div>
	);
};
