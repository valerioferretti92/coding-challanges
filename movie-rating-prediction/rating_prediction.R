#### Kaggle in Class Problem. ....
####Reference: https://inclass.kaggle.com/c/predict-movie-ratings
 
# Set data path as per your data file (for example: "c://abc//" )
setwd("E:/Universita/Mineria de Datos Aplicada/Trabajo/NUOVO_TRABAJO")
 
# If not installed, first install following three packages in R
install.packages('recommenderlab')
install.packages('reshape2')
install.packages('ggplot2')
library(recommenderlab)
library(reshape2)
library(ggplot2)
# Read training file along with header
tr<-read.csv("train.csv",header=TRUE)
# Just look at first few lines of this file
head(tr)
# Remove 'id' column. We do not need it
tr<-tr[,-c(1)]
# Check, if removed
tr[tr$user==1,]
# Using acast to convert above data as follows:
#       m1  m2   m3   m4
# u1    3   4    2    5
# u2    1   6    5
# u3    4   4    2    5
g<-acast(tr, user ~ movie)
g[1,1:100]
 
# Convert it as a matrix
R<-as.matrix(g)
 
# Convert R into realRatingMatrix data structure
#   realRatingMatrix is a recommenderlab sparse-matrix like data-structure
r <- as(R, "realRatingMatrix")
r
 
# view r in other possible ways
as(r, "list")     # A list
as(r, "matrix")   # A sparse matrix
 
# I can turn it into data-frame
head(as(r, "data.frame"))
 
# Draw an image plot of raw-ratings & normalized ratings
#  A column represents one specific movie and ratings by users
#   are shaded.
#   Note that some items are always rated 'black' by most users
#    while some items are not rated by many users
#     On the other hand a few users always give high ratings
#      as in some cases a series of black dots cut across items
image(r, main = "Raw Ratings")
 
# Create a recommender object (model)
#   Run anyone of the following four code lines.
#     Do not run all four
#       They pertain to four different algorithms.
#        UBCF: User-based collaborative filtering
#        IBCF: Item-based collaborative filtering
#      Parameter 'method' decides similarity measure
#        Cosine or Jaccard
rec=Recommender(r[1:nrow(r)],method="UBCF", param=list(normalize = "Z-score",method="Cosine"))
rec=Recommender(r[1:nrow(r)],method="IBCF", param=list(normalize = "Z-score",method="Cosine"))
 
############Create predictions#############################
# This prediction does not predict movie ratings for test.
#   But it fills up the user 'X' item matrix so that
#    for any userid and movieid, I can find predicted rating
#     dim(r) shows there are 6040 users (rows)
#      'type' parameter decides whether you want ratings or top-n items
#         get top-10 recommendations for a user, as:
#             predict(rec, r[1:nrow(r)], type="topNList", n=10)
recom <- predict(rec, r[1:nrow(r)], type="ratings")
recom

 
########## Create submission File from model #######################
# Read test file
test<-read.csv("test.csv",header=TRUE)
head(test)
# Get ratings list
rec_list<-as(recom,"list")
head(summary(rec_list))
ratings<-NULL
# For all lines in test file, one by one
for ( u in 1:length(test[,2]))
{
   # Read userid and movieid from columns 2 and 3 of test data
   userid <- test[u,2]
   movieid<-test[u,3]
 
   # Get as list & then convert to data frame all recommendations for user: userid
   u1<-as.data.frame(rec_list[[userid]])
   # Create a (second column) column-id in the data-frame u1 and populate it with row-names
   # Remember (or check) that rownames of u1 contain are by movie-ids
   # We use row.names() function
   u1$id<-row.names(u1)
   # Now access movie ratings in column 1 of u1
   x= u1[u1$id==movieid,1]
   # print(u)
   # print(length(x))
   # If no ratings were found, assign 0. You could also
   #   assign user-average
   if (length(x)==0)
   {
     ratings[u] <- 2
   }
   else
   {
     ratings[u] <- x
   }
 
}
length(ratings)
tx<-cbind(test[,1],round(ratings))
# Write to a csv file: submitfile.csv in your folder
write.table(tx,file="submitfile_IBCF.csv",row.names=FALSE,col.names=FALSE,sep=',')
# Submit now this csv file to kaggle
########################################

########## Model Evaluation #######################
# create evaluation scheme splitting taking 90% of the date for training and leaving 10% for validation or test
e <- evaluationScheme(r[1:nrow(r)], method="split", train=0.9, given=9)

# creation of recommender model based on ubcf
Rec.ubcf <- Recommender(getData(e, "train"), method="UBCF", param=list(normalize = "Z-score",method="Cosine"))

# creation of recommender model based on ibcf for comparison
Rec.ibcf <- Recommender(getData(e, "train"), method="IBCF", param=list(normalize = "Z-score",method="Cosine"))

# making predictions on the test data set
p.ubcf <- predict(Rec.ubcf, getData(e, "known"), type="ratings")

# making predictions on the test data set
p.ibcf <- predict(Rec.ibcf, getData(e, "known"), type="ratings")

# obtaining the error metrics for both approaches and comparing them
error.ubcf<-calcPredictionAccuracy(p.ubcf, getData(e, "unknown"))
error.ibcf<-calcPredictionAccuracy(p.ibcf, getData(e, "unknown"))
error <- rbind(error.ubcf,error.ibcf)
rownames(error) <- c("UBCF","IBCF")
error
