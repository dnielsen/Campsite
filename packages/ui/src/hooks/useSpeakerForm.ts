import * as Yup from "yup";
import { useHistory } from "react-router-dom";
import {
  FormSpeakerInput,
  SpeakerPreview,
  UseForm,
} from "../common/interfaces";
import { BASE_SPEAKER_API_URL } from "../common/constants";

export default function useSpeakerForm(): UseForm<FormSpeakerInput> {
  const history = useHistory();

  async function onSubmit(input: FormSpeakerInput) {
    // Send a request to create the speaker.
    console.log(input);
    const createdSpeaker = (await fetch(BASE_SPEAKER_API_URL, {
      method: "POST",
      body: JSON.stringify(input),
    }).then((res) => res.json())) as SpeakerPreview;
    // Redirect to the created speaker page.
    history.push(`/speakers/${createdSpeaker.id}`);
  }

  const initialValues: FormSpeakerInput = {
    name: "",
    photo: "",
    headline: "",
    bio: "",
  };

  const validationSchema = Yup.object().shape({});

  const formConfig = {
    onSubmit,
    initialValues,
    validationSchema,
  };

  return { formConfig };
}