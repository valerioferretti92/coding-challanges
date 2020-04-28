# Movie Rating Prediction, Kaggle

The objective of this is to provide predictions about the evaluation and the level of user appreciation about movies they have not seen yet based on their ratings of previously viewed movies. The movies that make up our database have been downloaded from Kaggle: https://www.kaggle.com/c/predict-movie-ratings-v22/overview.  
Specifically, three files are provieded:  
 
- The file TRAIN_V2 contains 750156 lines each made by 4 attributes: ID, USER_ID, MOVIE_ID, RATING. The first three are
numbers are respectively an ID, which uniquly identifies a line of the file, a user ID and a movie ID. The last attribute is an evaluation of the movie where the minimum value is 1 and the maximum is 5. We will be using this file to train our user preferences prediction model.  
- The file TEST_V2 contains 250053 lines each one made by three attributes: ID, USER_ID, MOVIE_ID. Each line correspond to a pair user-movie such that that specific user has never seen the specific movie associeted to him in this file. Our aim will be provide a predicted rating for every line of this file using a model trained with the file TRAIN_V2
- The SampleSubmission file is a model of what should be the format of the file that we have to upload to Kaggle. This target file has two fields: an ID field and an evaluation between 1 and 5. The ID field refers to a line in the TEST_V2 file. The evaluation field would be the level of appreciation expected for the user USER_ID on the movie MOVIE_ID in the file TEST_V2.

The file rating_prediction.R analyses data and produces the submission file with either UBCF or IBCF algorithms.
