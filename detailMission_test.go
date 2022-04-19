package main

import (
	"html"
	"testing"
)

var missionDetail *LbcDetail

func Test_missionDetail_Example1_ContactName_Shouldbe_Correct(t *testing.T) {
	missionDetail = new(LbcDetail)

	missionDetail.missionTextToDetailObjectConverter(missionExample1)

	if missionDetail.ContactName != "Steve AUSTIN" {
		t.Errorf("ContactName should be 'Steve AUSTIN' but was '%s'",
			missionDetail.ContactName)
	}
}

func Test_missionDetail_Example1_Description_Shouldbe_Correct(t *testing.T) {
	missionDetail = new(LbcDetail)

	expected := `Tripatols accélère sa transformation digitale avec pour mission de devenir LA plateforme numérique sportive qui permettra aux clients de découvrir notre univers à travers de nombreuses expériences locales centrées sur le sport.`

	missionDetail.missionTextToDetailObjectConverter(missionExample1)

	if missionDetail.Description != expected {
		t.Errorf("Description should be '%s' but was '%s'",
			expected,
			missionDetail.Description)
	}
}

func Test_missionDetail_Example1_Skills_Shouldbe_Correct(t *testing.T) {
	missionDetail = new(LbcDetail)

	expected := `* Agile : Expert * Kotlin : Expert`

	missionDetail.missionTextToDetailObjectConverter(missionExample1)

	if missionDetail.Skills != expected {
		t.Errorf("Skills should be '%s' but was '%s'",
			html.EscapeString(expected),
			html.EscapeString(missionDetail.Skills))
	}
}

func Test_missionDetail_Example2_Skills_Shouldbe_Correct(t *testing.T) {
	missionDetail = new(LbcDetail)

	expected := `* SQL : Confirmé * Talend : Débutant * Opcon : Débutant * ITIL : Débutant`

	missionDetail.missionTextToDetailObjectConverter(missionExample2)

	if missionDetail.Skills != expected {
		t.Errorf("Skills should be '%s' but was '%s'",
			expected,
			missionDetail.Skills)
	}
}

func Test_missionDetail_Example1_Goals_Shouldbe_Correct(t *testing.T) {
	missionDetail = new(LbcDetail)

	expected := `* Tu as au moins 4 ans d’expérience dans le développement Android.  * Tu maîtrises git et GitHub.`

	missionDetail.missionTextToDetailObjectConverter(missionExample1)

	if missionDetail.Goals != expected {
		t.Errorf("Goals should be '%s' but was '%s'",
			html.EscapeString(expected),
			html.EscapeString(missionDetail.Goals))
	}
}

func Test_missionDetail_Example1_Deliverables_Shouldbe_Correct(t *testing.T) {
	missionDetail = new(LbcDetail)

	expected := `Développeur Android confirmé/expert  * Début souhaité : 08/04/2022 * Date de fin du BDC : 31/12/2022`

	missionDetail.missionTextToDetailObjectConverter(missionExample1)

	if missionDetail.Deliverables != expected {
		t.Errorf("Deliverables should be '%s' but was '%s'",
			html.EscapeString(expected),
			html.EscapeString(missionDetail.Deliverables))
	}
}

var missionExample1 string = `<iframe src="https://www.googletagmanager.com/ns.html?id=" + GTM-5GD7MJP height="0" width="0" style="display:none;visibility:hidden"></iframe>

* Appels d'Offres ( /v2/supplier/request-for-proposal/search )
* Développeur Android confirmé/expert pour Tripatols ( /rfp/4uGHDw15mSeEWACKAFABCD/ )

**************************************************
Développeur Android confirmé/expert pour Tripatols
RFP-88036-1

**************************************************

Répondre à l'Appel d'Offres ( /presta/rfp/4uGHDw15mSeEWACKAFABCD/apply/#form )

* Resp. Op. : Steve AUSTIN 2350 avis ( /company-overview/wFvqVQv-ONbOxn79iABCD/overview/ )
* Créé le : 07/04/2022
* Relancer les référencés ( javascript: )

Ajouter aux favoris ( javascript:; )

----------------------
Contexte de la mission
----------------------

Tripatols accélère sa transformation digitale avec pour mission de devenir LA plateforme numérique sportive qui permettra aux clients de découvrir notre univers à travers de nombreuses expériences locales centrées sur le sport.

---------------------------------
Objectifs de la mission et tâches
---------------------------------

* Tu as au moins 4 ans d’expérience dans le développement Android.

* Tu maîtrises git et GitHub.

---------
Livrables
---------

Développeur Android confirmé/expert

* Début souhaité : 08/04/2022
* Date de fin du BDC : 31/12/2022

-------------------------
Compétences informatiques
-------------------------

* Agile : Expert
* Kotlin : Expert

------------------------
Adresse de réalisation :
------------------------

( https://maps.google.com/maps?q=Lille%2C%20France )
42 rue du pont, 59000 Lille, France

--------------
Déplacements :
--------------

42 rue du pont, 59000 Lille, France ( https://maps.google.com/maps?q=Lille%2C%20France ) (3/semaine)
Voir plus Voir moins ( javascript:; )

------------------------------
Questions / Réponses publiques
------------------------------

( /my/resume/ )
Exprimez-vous Partager ( javascript: )`

var missionExample2 string = `<iframe src="https://www.googletagmanager.com/ns.html?id=" + GTM-5GD7MJP height="0" width="0" style="display:none;visibility:hidden"></iframe>

* Appels d'Offres ( /v2/supplier/request-for-proposal/search )
* Ingénieur Support Data Confirmé pour Tripatols ( /rfp/j4ywOAMD4sSGsR-Uj7ABCD/ )

**********************************************
Ingénieur Support Data Confirmé pour Tripatols

**********************************************

Répondre à l'Appel d'Offres ( /presta/rfp/j4ywOAMD4sSGsR-Uj7ABCD/apply/#form )

* Resp. Op. : Steve AUSTIN 2350 avis ( /company-overview/wFvqVQv-ONbOxn79irABCD/overview/ )
* Créé le : 03/03/2022 | Modifié le : 07/04/2022
* Relancer les référencés ( javascript: )

Ajouter aux favoris ( javascript:; )

---------------------------------
Objectifs de la mission et tâches
---------------------------------

* **Diffuser les consignes de mise en production et de qualité de données.**

* **Assurer un support aux utilisateurs dans un contexte DATA (analyse de données)**

* *Participer à la mise en place et diffuser les indicateurs (performance, coût, délai, méthode SRE), les normes (ITIL) et les procédures d’exploitation.*

* *Définir/améliorer les systèmes d’alerting en production*

* **Suivre les incidents d’exploitation, en analyser les causes et prendre les mesures correctives associées (rédaction de post-mortem)**

* Intervenir sous contraintes de délais lors des incidents d’exploitation (analyse des incidents, diagnostic et résolution des incidents).

* Réaliser des astreintes dans une rotation d'équipe (1 weekend/mois)

* Anticiper les besoins des utilisateurs et les évolutions du système.

* Établir la documentation

---------
Livrables
---------

Atteinte des KPIs du support

* Début souhaité : 04/03/2022
* Date de fin du BDC : 02/05/2022
* Durée du BDC : 30 jours
* Durée estimée de la mission :
600 jours

* 
*Villeneuve-d'Ascq, France - En présentiel*

* ! Montant incluant les frais d'utilisation de la plateforme

Le client a demandé la saisie obligatoire des compétences de l'Appel d'Offre

-------------------------
Compétences informatiques
-------------------------

* SQL : Confirmé
* Talend : Débutant
* Opcon : Débutant
* ITIL : Débutant

--------------------------
Compétences fonctionnelles
--------------------------

* Gestion de crise : Confirmé
* Connaissance milieu Retail/Encaissement : Débutant

-------
Langues
-------

* français : Courant
* Anglais / English : Intermédiaire

------------------------
Adresse de réalisation :
------------------------

( https://maps.google.com/maps?q=Lille%2C%20France )
42 rue du pont, 59000 Lille, France

--------------
Déplacements :
--------------

42 rue du pont, 59000 Lille, France ( https://maps.google.com/maps?q=Lille%2C%20France ) (2 jours/semaine)
Voir plus Voir moins ( javascript:; )

------------------------------
Questions / Réponses publiques
------------------------------

( /my/resume/ )
Exprimez-vous Partager ( javascript: )`
