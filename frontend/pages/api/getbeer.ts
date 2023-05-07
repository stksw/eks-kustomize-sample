// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import axios, { AxiosResponse } from 'axios';
import { NextApiResponse, NextApiRequest } from 'next/types';

type Data = {
  message: string;
  beer: any;
};

export default function handler(req: NextApiRequest, res: NextApiResponse<Data>) {
  const BACKEND_URL = process.env.ON_EKS
    ? 'http://backend.eks-sample-backend/api/getbeer'
    : 'http://backend-service:8080/api/getbeer';

  axios
    .get(BACKEND_URL)
    .then((response: AxiosResponse<string>) => {
      console.log(response);
      res.status(200).json({ message: 'Success!!', beer: response.data });
    })
    .catch((err) => {
      console.log(err);
      res.status(400).json({ message: 'Bad Request!', beer: '' });
    });
}
