package repositories

import (
	"fmt"
	"hallocorona/models"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	FindArticles() ([]models.Article, error)
	GetArticle(ID int) (models.Article, error)
	CreateArticle(article models.Article) (models.Article, error)
	UpdateArticle(article models.Article) (models.Article, error)
	UpdateArticleCategory(article models.Article, articleId int, categoryId []int) (models.Article, error)
	DeleteArticle(article models.Article) (models.Article, error)
	GetArticleAuthor(UserID int) (models.User, error)
	FindArticleCategoriesByID(CategoryID []int) ([]models.Category, error)
}

func RepositoryArticle(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindArticleCategoriesByID(CategoryID []int) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories, CategoryID).Error

	return categories, err
}

func (r *repository) GetArticleAuthor(UserID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, UserID).Error

	return user, err
}

func (r *repository) FindArticles() ([]models.Article, error) {
	var articles []models.Article
	err := r.db.Preload("User").Preload("Category").Find(&articles).Error

	return articles, err
}

func (r *repository) GetArticle(ID int) (models.Article, error) {
	var article models.Article
	err := r.db.Preload("User").Preload("Category").First(&article, ID).Error

	return article, err
}

func (r *repository) CreateArticle(article models.Article) (models.Article, error) {
	err := r.db.Preload("User").Preload("Category").Create(&article).Error

	return article, err
}

func (r *repository) UpdateArticle(article models.Article) (models.Article, error) {
	err := r.db.Preload("User").Preload("Category").Save(&article).Error

	return article, err
}

func (r *repository) UpdateArticleCategory(article models.Article, articleId int, categoryId []int) (models.Article, error) {
	// delete semua data terkait dulu
	fmt.Println("ini category id bro", categoryId)
	var err error
	err = r.db.Exec("DELETE FROM article_categories WHERE article_id=?", articleId).Error
	for _, id := range categoryId {
		fmt.Println("ini idnyaaa ", id)
		// kenapa gini, karena gatau kenapa ada angka 43 45, gaje banget emang golang T_T
		if id == 43 || id == 45 {
			continue
		}
		err = r.db.Exec("INSERT INTO article_categories VALUES (?,?)", articleId, id).Error
	}

	return article, err
}

func (r *repository) DeleteArticle(article models.Article) (models.Article, error) {
	err := r.db.Preload("User").Preload("Category").Delete(&article).Error

	return article, err
}
