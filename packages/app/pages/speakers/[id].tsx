import { GetStaticProps, GetStaticPaths } from "next";
import Layout from "../../components/Layout";
import React from "react";
import { Speaker, SpeakerPreview } from "../../common/interfaces";
import { BASE_API_URL } from "../../common/constants";
import SpeakerItem from "../../components/speaker/SpeakerItem";

const BASE_SPEAKER_API_URL = `${BASE_API_URL}/speakers`;

interface Props {
  data: any;
}

const StaticPropsDetail = (props: Props) => {
  return (
    <Layout title={`${props.data.name}`}>
      <SpeakerItem speaker={props.data} />
    </Layout>
  );
};

export default StaticPropsDetail;

export const getStaticProps: GetStaticProps = async ({
  params,
}): Promise<{ props: Props }> => {
  const res = await fetch(`${BASE_SPEAKER_API_URL}/${params?.id}`);
  const speaker: Speaker = await res.json();
  return { props: { data: speaker } };
};

export const getStaticPaths: GetStaticPaths = async () => {
  const res = await fetch(BASE_SPEAKER_API_URL);
  const speakers: Speaker[] = await res.json();
  const paths = speakers.map((s) => ({ params: { id: s.id } }));
  return {
    paths,
    fallback: false,
  };
};
